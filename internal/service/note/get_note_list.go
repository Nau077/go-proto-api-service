package note

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (s *Service) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {
	res, err := s.noteRepository.GetNoteList(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
