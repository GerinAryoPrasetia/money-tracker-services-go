-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    email,
    password,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id;