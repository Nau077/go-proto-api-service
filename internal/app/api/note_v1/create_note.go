package note_v1

import (
	"context"

	desc "github.com/Nau077/golang-pet-first/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := n.noteService.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
	// dbDsn := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	// 	host, port, dbUser, dbPassword, dbName, sslMode,
	// )

	// db, err := sqlx.Open("pgx", dbDsn)
	// if err != nil {
	// 	return nil, err
	// }
	// defer db.Close()

	// builder := sq.Insert(noteTable).
	// 	PlaceholderFormat(sq.Dollar).
	// 	Columns("title, text, author").
	// 	Values(req.GetNoteContent().GetTitle(), req.GetNoteContent().GetText(), req.GetNoteContent().GetAuthor()).
	// 	Suffix("returning id")

	// query, args, err := builder.ToSql()
	// if err != nil {
	// 	return nil, err
	// }

	// row, err := db.QueryContext(ctx, query, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer row.Close()

	// row.Next()
	// var id int64
	// err = row.Scan(&id)
	// if err != nil {
	// 	return nil, err
	// }

	// return &desc.CreateNoteResponse{
	// 	Id: id,
	// }, nil
}
