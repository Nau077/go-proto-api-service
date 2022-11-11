package note

import (
	"context"
)

func (s *Service) DeleteNote(ctx context.Context, noteId int64) error {
	return s.noteRepository.DeleteNote(ctx, noteId)
}
