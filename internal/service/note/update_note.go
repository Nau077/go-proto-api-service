package note

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	res, err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
