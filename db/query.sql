-- name: CreateNotebook :exec
INSERT INTO notebooks (
    alias_id, ulid, name
) VALUES (
    ?, ?, ?
);

-- name: GetNotebook :one
SELECT *
FROM notebooks
WHERE alias_id = ?;

-- name: GetBelongNotebook :many
SELECT *
FROM notebooks_users AS NU
INNER JOIN users AS U
ON NU.user_id = U.ulid
WHERE U.uid = ?;

-- name: AddUserToNotebook :exec
INSERT INTO notebooks_users (
    notebook_id, user_id
) VALUES (
    ?, ?
);

-- name: AddUserToNotebookAsAdmin :exec
INSERT INTO notebooks_users (
    notebook_id, user_id, is_admin
) VALUES (
    ?, ?, 1
);

-- name: CreateShout :exec
INSERT INTO shouts (
    ulid, notebook_id, user_id
) VALUES (
    ?, ?, ?
);

-- name: GetShout :one
SELECT *
FROM shouts
WHERE ulid = ?;

-- name: CreateUser :exec
INSERT INTO users (
    ulid, uid, alias_id, name
) VALUES (
    ?, ?, ?, ?
);

-- name: GetUserID :one
SELECT *
FROM users
WHERE uid = ?;

-- name: DeleteShout :exec
UPDATE shouts
SET deleted_at = ?
WHERE ulid = ?;
