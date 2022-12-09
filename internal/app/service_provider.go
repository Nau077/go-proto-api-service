package app

import (
	"context"
	"log"

	"github.com/Nau077/golang-pet-first/internal/config"
	"github.com/Nau077/golang-pet-first/internal/pkg/db"
	repository "github.com/Nau077/golang-pet-first/internal/repository/note"
	"github.com/Nau077/golang-pet-first/internal/service/note"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.Config
	// repositories
	noteRepository repository.Repository
	// services
	noteService *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetDB
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("cant connect to db err %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

// GetConfig
func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}
		s.config = cfg
	}

	return s.config
}

// GetNoteRepository
func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.Repository {
	if s.noteRepository == nil {
		s.noteRepository = repository.NewNoteRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

// GetNoteService
func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
