package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("UpdatetNote")

	return &desc.UpdateNoteResponse{
		Id: 3,
	}, nil
}
