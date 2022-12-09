package note

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/model"
)

func (s *Service) UpdateNote(ctx context.Context, req *model.UpdateNoteInfo) (int64, error) {
	res, err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return 0, err
	}

	return res, nil
}
