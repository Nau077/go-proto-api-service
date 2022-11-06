package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (n *Note) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {

	res, err := n.noteService.GetNoteList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
