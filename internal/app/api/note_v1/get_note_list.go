package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type record struct {
	id     int64
	title  string
	text   string
	author string
}

func (n *Note) GetNoteList(ctx context.Context, req *desc.Empty) (*desc.GetNoteListResponse, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select(noteTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
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
