package note_v1

import desc "github.com/Nau077/golang-pet-first/pkg/note_v1"

type Note struct {
	desc.UnimplementedNoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}
