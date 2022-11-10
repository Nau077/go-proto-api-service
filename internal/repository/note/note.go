package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Nau077/golang-pet-first/internal/repository/table"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	DeleteNote(ctx context.Context, userId int64) error
	GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error)
}

type repository struct {
	db *sqlx.DB
}

type record struct {
	id        int64
	title     string
	text      string
	author    string
	email     string
	createdAt time.Time
	updatedAt *time.Time
}

func NewNoteRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author, email").
		Values(req.GetNoteContent().GetTitle(), req.GetNoteContent().GetText(), req.GetNoteContent().GetEmail(), req.GetNoteContent().GetAuthor()).
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

func (r *repository) DeleteNote(ctx context.Context, userId int64) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userId})

	_, _, err := builder.ToSql()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {
	builder := sq.Select("id, title, author, text, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note)

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
		err := rows.Scan(&rec.id, &rec.title, &rec.text, &rec.author, &rec.email, &rec.createdAt, &rec.updatedAt)
		if err != nil {
			return nil, err
		}
		recordsList = append(recordsList, rec)
	}

	var recordsListProto []*desc.Record

	for _, v := range recordsList {
		// создаём пустой указатель
		var updatedAt *timestamppb.Timestamp
		if v.updatedAt != nil {
			updatedAt = timestamppb.New(*v.updatedAt)
		}

		recordsListProto = append(recordsListProto,
			&desc.Record{
				Id:        v.id,
				CreatedAt: timestamppb.New(v.createdAt),
				UpdatedAt: updatedAt,
				NoteContent: &desc.NoteContent{
					Title:  v.title,
					Text:   v.text,
					Author: v.author,
					Email:  v.email,
				},
			})
	}

	return &desc.GetNoteListResponse{
		Records: recordsListProto,
	}, nil
}

func (r repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	builder := sq.Select("id, title, author, text, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

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
	var email string
	var createdAt time.Time
	var updatedAt *time.Time

	err = row.Scan(&id, &title, &author, &text, &email, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	var updatedAtMpb *timestamppb.Timestamp
	if updatedAt != nil {
		updatedAtMpb = timestamppb.New(*updatedAt)
	}

	return &desc.GetNoteResponse{
		Record: &desc.Record{
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtMpb,
			NoteContent: &desc.NoteContent{
				Title:  title,
				Author: author,
				Text:   text,
				Email:  email,
			},
			Id: id,
		},
	}, nil
}

func (r repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("title", req.GetNoteContent().GetTitle()).
		Set("text", req.GetNoteContent().GetText()).
		Set("author", req.GetNoteContent().GetAuthor()).
		Set("email", req.GetNoteContent().GetEmail()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()}).
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
