package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")

	return &desc.GetNoteResponse{
		Id: 4,
	}, nil
}
