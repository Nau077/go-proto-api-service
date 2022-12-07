package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Nau077/golang-pet-first/internal/app/api/note_v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
)

// App
type App struct {
	noteImpl        *note_v1.Note
	serviceProvider *serviceProvider
	pathConfig      string
	grpcServer      *grpc.Server
	mux             *runtime.ServeMux
}

// NewApp
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

// Run
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	// nolint:staticcheck
	opts := []grpc.DialOption{grpc.WithInsecure}

	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().GRPC.GetAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.GetAddress())
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process grpc server: #{err.Error}")
		}
	}()

	log.Printf("run grpc server on host %s", a.serviceProvider.GetConfig().GRPC.GetAddress())

	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(a.serviceProvider.GetConfig().HTTP.GetAddress(), a.mux); err != nil {
			log.Fatalf("failed to process muxer: #{err.Error()}")
		}
	}()

	log.Printf("run http server on host %s", a.serviceProvider.GetConfig().HTTP.GetAddress())

	return nil
}