package note_v1

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/app/converter"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	id, err := n.noteService.UpdateNote(ctx, converter.ToUpdateNoteInfo(req.GetUpdateNoteInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateNoteResponse{
		Id: id,
	}, nil
}
