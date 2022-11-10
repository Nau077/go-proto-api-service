package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Nau077/golang-pet-first/internal/app/api/note_v1"
	repository "github.com/Nau077/golang-pet-first/internal/repository/note"
	"github.com/Nau077/golang-pet-first/internal/service/note"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const (
	hostGrpc = ":50051"
	hostHttp = ":8090"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		// nolint:errcheck
		error := startGRPC()
		if error != nil {
			log.Fatalf("failed running grpc server: %s", error.Error())
		}
	}()

	go func() {
		defer wg.Done()
		error := startHttp()
		if error != nil {
			log.Fatalf("failed running http server: %s", error.Error())
		}
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port %s", err.Error())
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)

	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteServiceServer(s, note_v1.NewNote(noteService))

	fmt.Println("grpc Server is running on port:", hostGrpc)

	if err = s.Serve(list); err != nil {
		return err
	}

	return nil
}

func startHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// nolint:staticcheck
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	fmt.Println("http Server is running on port:", hostHttp)

	if err = http.ListenAndServe(hostHttp, mux); err != nil {
		return err
	}
	return nil
}
