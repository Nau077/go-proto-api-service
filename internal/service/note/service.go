package note

import repository "github.com/Nau077/golang-pet-first/internal/repository/note"

type Service struct {
	noteRepository repository.Repository
}

func NewService(noteRepository repository.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
