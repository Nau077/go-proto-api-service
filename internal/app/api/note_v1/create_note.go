package note_v1

import (
	"context"
	"fmt"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title", req.GetTitle())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
