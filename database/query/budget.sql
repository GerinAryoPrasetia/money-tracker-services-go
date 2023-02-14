-- name: CreateBudget :one
INSERT INTO budget (
    id,
    user_id,
    amount,
    description,
    created_by
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBudgetByUserID :one
SELECT * FROM budget
WHERE user_id = $1 AND deleted_at <> null LIMIT 1;

-- name: GetListBudget :many
SELECT * FROM budget
WHERE user_id = $1
LIMIT $1
OFFSET $2;

-- name: DeleteBudget :exec
UPDATE budget
SET
deleted_at = NOW(),
deleted_by = $1
WHERE user_id = $2;