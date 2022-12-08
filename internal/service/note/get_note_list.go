package note

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/model"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (s *Service) GetNoteList(ctx context.Context, req *desc.Empty) (*[]model.Record, error) {
	records, err := s.noteRepository.GetNoteList(ctx, req)
	if err != nil {
		return &[]model.Record{}, err
	}

	return records, nil
}
