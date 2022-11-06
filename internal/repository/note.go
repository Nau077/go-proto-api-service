package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Nau077/golang-pet-first/internal/repository/table"
	_ "github.com/Nau077/golang-pet-first/internal/repository/table"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetNoteContent().GetTitle(), req.GetNoteContent().GetText(), req.GetNoteContent().GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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
