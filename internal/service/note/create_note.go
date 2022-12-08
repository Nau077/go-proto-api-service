package note

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/model"
)

func (s *Service) CreateNote(ctx context.Context, noteContent *model.NoteContent) (int64, error) {
	id, err := s.noteRepository.CreateNote(ctx, noteContent)
	if err != nil {
		return 0, err
	}

	return id, nil
}
