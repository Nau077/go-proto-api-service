package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Nau077/golang-pet-first/internal/model"
	"github.com/Nau077/golang-pet-first/internal/pkg/db"
	"github.com/Nau077/golang-pet-first/internal/repository/table"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
)

type Repository interface {
	CreateNote(ctx context.Context, noteContent *model.NoteContent) (int64, error)
	DeleteNote(ctx context.Context, userId int64) error
	GetNoteList(ctx context.Context, req *desc.Empty) ([]*model.Record, error)
	GetNote(ctx context.Context, id int64) (*model.Record, error)
	UpdateNote(ctx context.Context, updateNoteInfo *model.UpdateNoteInfo) (int64, error)
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateNote(ctx context.Context, noteContent *model.NoteContent) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author, email").
		Values(noteContent.Title, noteContent.Text, noteContent.Email, noteContent.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}
	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteNote(ctx context.Context, userId int64) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}
	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetNoteList(ctx context.Context, req *desc.Empty) ([]*model.Record, error) {
	builder := sq.Select("id, title, author, text, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "GetNoteList",
		QueryRaw: query,
	}
	var records []*model.Record

	err = r.client.DB().SelectContext(ctx, records, q, args...)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repository) GetNote(ctx context.Context, id int64) (*model.Record, error) {
	builder := sq.Select("id, title, author, text, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}
	var note = new(model.Record)
	err = r.client.DB().GetContext(ctx, note, q, args...)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *repository) UpdateNote(ctx context.Context, updateNoteInfo *model.UpdateNoteInfo) (int64, error) {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now()).
		Suffix("returning id")

	if updateNoteInfo.Title.Valid {
		builder.Set("title", updateNoteInfo.Title.String)
	}

	if updateNoteInfo.Text.Valid {
		builder.Set("title", updateNoteInfo.Text.String)
	}

	if updateNoteInfo.Author.Valid {
		builder.Set("title", updateNoteInfo.Author.String)
	}

	if updateNoteInfo.Email.Valid {
		builder.Set("title", updateNoteInfo.Email.String)
	}

	builder = builder.Where(sq.Eq{"id": updateNoteInfo.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Update",
		QueryRaw: query,
	}

	_, err = r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return updateNoteInfo.Id, nil
}
