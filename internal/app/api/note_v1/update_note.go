package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	res, err := n.noteService.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
