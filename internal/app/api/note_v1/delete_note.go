package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	fmt.Println("DeleteNote")

	return &desc.DeleteNoteResponse{
		Id: 2,
	}, nil
}
