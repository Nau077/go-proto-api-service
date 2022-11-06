package note_v1

import (
	"github.com/Nau077/golang-pet-first/internal/service/note"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteServiceServer

	noteService *note.Service
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}
