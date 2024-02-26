// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUserToNotebook = `-- name: AddUserToNotebook :exec
INSERT INTO notebooks_users (
    notebook_id, user_id
) VALUES (
    $1, $2
)
`

type AddUserToNotebookParams struct {
	NotebookID string
	UserID     string
}

func (q *Queries) AddUserToNotebook(ctx context.Context, arg AddUserToNotebookParams) error {
	_, err := q.db.Exec(ctx, addUserToNotebook, arg.NotebookID, arg.UserID)
	return err
}

const addUserToNotebookAsAdmin = `-- name: AddUserToNotebookAsAdmin :exec
INSERT INTO notebooks_users (
    notebook_id, user_id, is_admin
) VALUES (
    $1, $2, 1
)
`

type AddUserToNotebookAsAdminParams struct {
	NotebookID string
	UserID     string
}

func (q *Queries) AddUserToNotebookAsAdmin(ctx context.Context, arg AddUserToNotebookAsAdminParams) error {
	_, err := q.db.Exec(ctx, addUserToNotebookAsAdmin, arg.NotebookID, arg.UserID)
	return err
}

const createNotebook = `-- name: CreateNotebook :exec
INSERT INTO notebooks (
    alias_id, ulid, name
) VALUES (
    $1, $2, $3
)
`

type CreateNotebookParams struct {
	AliasID string
	Ulid    string
	Name    string
}

func (q *Queries) CreateNotebook(ctx context.Context, arg CreateNotebookParams) error {
	_, err := q.db.Exec(ctx, createNotebook, arg.AliasID, arg.Ulid, arg.Name)
	return err
}

const createShout = `-- name: CreateShout :exec
INSERT INTO shouts (
    ulid, notebook_id, user_id
) VALUES (
    $1, $2, $3
)
`

type CreateShoutParams struct {
	Ulid       string
	NotebookID string
	UserID     string
}

func (q *Queries) CreateShout(ctx context.Context, arg CreateShoutParams) error {
	_, err := q.db.Exec(ctx, createShout, arg.Ulid, arg.NotebookID, arg.UserID)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    ulid, uid, alias_id, name
) VALUES (
    $1, $2, $3, $4
)
`

type CreateUserParams struct {
	Ulid    string
	Uid     string
	AliasID string
	Name    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Ulid,
		arg.Uid,
		arg.AliasID,
		arg.Name,
	)
	return err
}

const deleteShout = `-- name: DeleteShout :exec
UPDATE shouts
SET deleted_at = $1
WHERE ulid = $2
`

type DeleteShoutParams struct {
	DeletedAt pgtype.Timestamptz
	Ulid      string
}

func (q *Queries) DeleteShout(ctx context.Context, arg DeleteShoutParams) error {
	_, err := q.db.Exec(ctx, deleteShout, arg.DeletedAt, arg.Ulid)
	return err
}

const getBelongNotebook = `-- name: GetBelongNotebook :many
SELECT notebook_id, user_id, is_admin, ulid, uid, alias_id, name, created_at, updated_at, deleted_at
FROM notebooks_users AS NU
INNER JOIN users AS U
ON NU.user_id = U.ulid
WHERE U.uid = $1
`

type GetBelongNotebookRow struct {
	NotebookID string
	UserID     string
	IsAdmin    pgtype.Bool
	Ulid       string
	Uid        string
	AliasID    string
	Name       string
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
	DeletedAt  pgtype.Timestamptz
}

func (q *Queries) GetBelongNotebook(ctx context.Context, uid string) ([]GetBelongNotebookRow, error) {
	rows, err := q.db.Query(ctx, getBelongNotebook, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBelongNotebookRow
	for rows.Next() {
		var i GetBelongNotebookRow
		if err := rows.Scan(
			&i.NotebookID,
			&i.UserID,
			&i.IsAdmin,
			&i.Ulid,
			&i.Uid,
			&i.AliasID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNotebook = `-- name: GetNotebook :one
SELECT ulid, alias_id, name, created_at, updated_at, deleted_at
FROM notebooks
WHERE alias_id = $1
`

func (q *Queries) GetNotebook(ctx context.Context, aliasID string) (Notebook, error) {
	row := q.db.QueryRow(ctx, getNotebook, aliasID)
	var i Notebook
	err := row.Scan(
		&i.Ulid,
		&i.AliasID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getShout = `-- name: GetShout :one
SELECT ulid, name, notebook_id, user_id, script, created_at, updated_at, deleted_at
FROM shouts
WHERE ulid = $1
`

func (q *Queries) GetShout(ctx context.Context, ulid string) (Shout, error) {
	row := q.db.QueryRow(ctx, getShout, ulid)
	var i Shout
	err := row.Scan(
		&i.Ulid,
		&i.Name,
		&i.NotebookID,
		&i.UserID,
		&i.Script,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserID = `-- name: GetUserID :one
SELECT ulid, uid, alias_id, name, created_at, updated_at, deleted_at
FROM users
WHERE uid = $1
`

func (q *Queries) GetUserID(ctx context.Context, uid string) (User, error) {
	row := q.db.QueryRow(ctx, getUserID, uid)
	var i User
	err := row.Scan(
		&i.Ulid,
		&i.Uid,
		&i.AliasID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}