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
		Id:     updateNoteInfo.Id,
		Title:  sql.NullString{String: updateNoteInfo.GetTitle().GetValue(), Valid: true},
		Author: sql.NullString{String: updateNoteInfo.GetAuthor().GetValue(), Valid: true},
		Text:   sql.NullString{String: updateNoteInfo.GetText().GetValue(), Valid: true},
		Email:  sql.NullString{String: updateNoteInfo.GetEmail().GetValue(), Valid: true},
	}
}

func ToRecord(record *desc.Record) *model.Record {

	return &model.Record{
		Id:          record.GetId(),
		NoteContent: ToNoteContent(record.NoteContent),
		CreatedAt:   record.GetCreatedAt().AsTime(),
		UpdatedAt:   sql.NullTime{Time: record.GetUpdatedAt().AsTime(), Valid: true},
	}
}

func ToDeskRecord(record *model.Record) *desc.Record {
	var updatedAt *timestamppb.Timestamp
	if record.UpdatedAt.Valid {
		updatedAt = timestamppb.New(record.UpdatedAt.Time)
	}

	return &desc.Record{
		Id:          record.Id,
		NoteContent: ToDescNoteContent(record.NoteContent),
		CreatedAt:   timestamppb.New(record.CreatedAt),
		UpdatedAt:   updatedAt,
	}
}

func ToDescRecordsList(records *[]model.Record) []*desc.Record {
	var recordsList []*desc.Record

	for _, record := range *records {
		recordsList = append(recordsList, ToDeskRecord(&record))
	}
	return recordsList
}
