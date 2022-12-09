package note

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.Record, error) {
	res, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
