// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Notebook struct {
	Ulid      string
	AliasID   string
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz
}

type NotebooksUser struct {
	NotebookID string
	UserID     string
	IsAdmin    pgtype.Bool
}

type Shout struct {
	Ulid       string
	Name       string
	NotebookID string
	UserID     string
	Script     pgtype.Text
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
	DeletedAt  pgtype.Timestamptz
}

type User struct {
	Ulid      string
	Uid       string
	AliasID   string
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz
}
