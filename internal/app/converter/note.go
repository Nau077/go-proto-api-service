package converter

import (
	"database/sql"

	"github.com/Nau077/golang-pet-first/internal/model"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteContent(noteContent *desc.NoteContent) *model.NoteContent {
	return &model.NoteContent{
		Title:  noteContent.GetTitle(),
		Author: noteContent.GetAuthor(),
		Text:   noteContent.GetText(),
		Email:  noteContent.GetEmail(),
	}
}

func ToDescNoteContent(noteContent *model.NoteContent) *desc.NoteContent {
	return &desc.NoteContent{
		Title:  noteContent.Title,
		Author: noteContent.Author,
		Text:   noteContent.Text,
		Email:  noteContent.Email,
	}
}

func ToUpdateNoteInfo(updateNoteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	return &model.UpdateNoteInfo{
		Title:  sql.NullString{String: updateNoteInfo.GetTitle().GetValue(), Valid: true},
		Author: sql.NullString{String: updateNoteInfo.GetAuthor().GetValue(), Valid: true},
		Text:   sql.NullString{String: updateNoteInfo.GetText().GetValue(), Valid: true},
		Email:  sql.NullString{String: updateNoteInfo.GetEmail().GetValue(), Valid: true},
	}
}

func ToDeskUpdateNoteInfo(updateNoteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	return &model.UpdateNoteInfo{
		Title:  sql.NullString{String: updateNoteInfo.GetTitle().GetValue(), Valid: true},
		Author: sql.NullString{String: updateNoteInfo.GetAuthor().GetValue(), Valid: true},
		Text:   sql.NullString{String: updateNoteInfo.GetText().GetValue(), Valid: true},
		Email:  sql.NullString{String: updateNoteInfo.GetEmail().GetValue(), Valid: true},
	}
}

func ToRecord(record *desc.Record) *model.Record {
	return &model.Record{
		ID:          record.GetId(),
		NoteContent: record.GetNoteContent(),
		CreatedAt:   record.GetCreatedAt().AsTime(),
		UpdatedAt:   record.GetUpdatedAt(),
	}
}

func ToDeskRecord(record *model.Record) *desc.Record {
	var updatedAt *timestamppb.Timestamp
	if record.UpdatedAt.Valid {
		updatedAt = timestamppb.New(record.UpdatedAt.Time)
	}

	return &desc.Record{
		Id:          record.ID,
		NoteContent: ToDescNoteContent(record.NoteContent),
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   updatedAt,
	}
}
