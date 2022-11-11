package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := n.noteService.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
