package note_v1

import (
	"context"

	"github.com/Nau077/golang-pet-first/internal/app/converter"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

func (n *Note) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {

	records, err := n.noteService.GetNoteList(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteListResponse{
		Records: converter.ToDescRecordsList(records),
	}, nil
}
