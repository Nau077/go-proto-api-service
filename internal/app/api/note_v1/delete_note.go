package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	err := n.noteService.DeleteNote(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
