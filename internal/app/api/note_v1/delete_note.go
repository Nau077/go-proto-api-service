package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	err := n.noteService.DeleteNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
