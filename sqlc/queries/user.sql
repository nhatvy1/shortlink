-- name: GetUser :one
SELECT id, email
FROM users
WHERE id = $1;