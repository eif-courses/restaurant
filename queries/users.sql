-- name: CreateUser :one
INSERT INTO users(email, password) VALUES ($1, $2)
RETURNING id, email, created_at;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT id, email, created_at FROM users WHERE id = $1;

-- name: CreateSession :one
INSERT INTO sessions(user_id, session_token, expires_at) VALUES ($1,$2, $3)
RETURNING *;

-- name: GetSessionByToken :one
SELECT
    s.id,
    s.user_id,
    s.session_token,
    s.expires_at,
    s.created_at,
    u.id as user_id,
    u.email,
    u.created_at as user_created_at
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.session_token = $1 AND s.expires_at > $2;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE session_token = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at < $1;

-- name: DeleteUserSessions :exec
DELETE FROM sessions WHERE user_id = $1;

-- name: GetUserSessions :many
SELECT * FROM sessions WHERE user_id = $1 AND expires_at > $2 ORDER BY created_at DESC;




