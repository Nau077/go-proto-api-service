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
		startGRPC()
	}()
	// nolint:errcheck
	go startHttp(&wg)
	// в wg.Wait будет висеть, пока счетчик не обнулится
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
		log.Fatalf("failed to serve %s", err.Error())
	}

	return nil
}

func startHttp(wg *sync.WaitGroup) error {
	defer wg.Done()

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

	return http.ListenAndServe(hostHttp, mux)
}
