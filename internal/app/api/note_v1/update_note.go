package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	res, err := n.noteService.UpdateNote(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
