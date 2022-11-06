package note

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := s.noteRepository.GetNote(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
