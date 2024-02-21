-- name: CreateShout :exec
INSERT INTO shouts (
    uid, ulid
) VALUES (
    ?, ?
);

-- name: DeleteShout :exec
UPDATE shouts
SET deleted_at = ?
WHERE ulid = ?;
