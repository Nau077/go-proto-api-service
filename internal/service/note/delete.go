package note

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	err := s.noteRepository.DeleteNote(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
