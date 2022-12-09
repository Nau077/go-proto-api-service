package note_v1

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/app/converter"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := n.noteService.CreateNote(ctx, converter.ToNoteContent(req.GetNoteContent()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{Id: id}, nil
}
