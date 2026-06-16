-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = ?;