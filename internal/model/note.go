package model

import (
	"database/sql"
	"time"
)

type NoteContent struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
	Email  string `db:"email"`
}

type UpdateNoteInfo struct {
	Id     int64          `db:"id"`
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
	Email  sql.NullString `db:"email"`
}

type Record struct {
	Id          int64        `db:"id"`
	NoteContent *NoteContent `db:""`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
