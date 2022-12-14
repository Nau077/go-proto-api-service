package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	res, err := n.noteService.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
