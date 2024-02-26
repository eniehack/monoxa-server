-- name: CreateNotebook :exec
INSERT INTO notebooks (
    alias_id, ulid, name
) VALUES (
    $1, $2, $3
);

-- name: GetNotebook :one
SELECT *
FROM notebooks
WHERE alias_id = $1;

-- name: GetBelongNotebook :many
SELECT *
FROM notebooks_users AS NU
INNER JOIN users AS U
ON NU.user_id = U.ulid
WHERE U.uid = $1;

-- name: AddUserToNotebook :exec
INSERT INTO notebooks_users (
    notebook_id, user_id
) VALUES (
    $1, $2
);

-- name: AddUserToNotebookAsAdmin :exec
INSERT INTO notebooks_users (
    notebook_id, user_id, is_admin
) VALUES (
    $1, $2, 1
);

-- name: CreateShout :exec
INSERT INTO shouts (
    ulid, notebook_id, user_id
) VALUES (
    $1, $2, $3
);

-- name: GetShout :one
SELECT *
FROM shouts
WHERE ulid = $1;

-- name: CreateUser :exec
INSERT INTO users (
    ulid, uid, alias_id, name
) VALUES (
    $1, $2, $3, $4
);

-- name: GetUserID :one
SELECT *
FROM users
WHERE uid = $1;

-- name: DeleteShout :exec
UPDATE shouts
SET deleted_at = $1
WHERE ulid = $2;
