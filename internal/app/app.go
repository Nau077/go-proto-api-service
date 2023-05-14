package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Nau077/golang-pet-first/internal/app/api/note_v1"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// App
type App struct {
	note            *note_v1.Note
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

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.note = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteServiceServer(a.grpcServer, a.note)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().GRPC.GetAddress(), opts)
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
