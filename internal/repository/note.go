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
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error
	GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error)
}

type repository struct {
	db *sqlx.DB
}

type record struct {
	id     int64
	title  string
	text   string
	author string
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

func (r *repository) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {
	builder := sq.Select(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recordsList []record

	for rows.Next() {
		var rec record
		err := rows.Scan(&rec.id, &rec.title, &rec.author, &rec.text)
		if err != nil {
			return nil, err
		}
		recordsList = append(recordsList, rec)
	}

	var recordsListProto []*desc.Record

	for _, v := range recordsList {
		recordsListProto = append(recordsListProto, &desc.Record{
			Id: v.id,
			NoteContent: &desc.NoteContent{
				Title:  v.title,
				Text:   v.text,
				Author: v.author,
			},
		})
	}

	return &desc.GetNoteListResponse{
		Record: recordsListProto,
	}, nil
}

func (r repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	builder := sq.Select(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Columns("title, text, author").
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	var title string
	var author string
	var text string
	err = row.Scan(&id, &title, &author, &text)
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		Record: &desc.Record{
			NoteContent: &desc.NoteContent{
				Title:  title,
				Text:   text,
				Author: author,
			},
			Id: id,
		},
	}, nil
}

func (r repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("title", req.Record.NoteContent.Title).
		Set("text", req.Record.NoteContent.Text).
		Set("author", req.Record.NoteContent.Author).
		Where(sq.Eq{"id": req.Record.Id}).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.UpdateNoteResponse{
		Id: id,
	}, nil
}
