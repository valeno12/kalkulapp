-- name: CreateSession :execresult
INSERT INTO sessions (name, code, created_by, max_participants, status) 
VALUES (?, ?, ?, ?, 'active');

-- name: CreateUser :execresult
INSERT INTO users (session_id, name) VALUES (?, ?);

-- name: GetSessionByCode :one
SELECT id, name, code, created_by, max_participants, status 
FROM sessions WHERE code = ? LIMIT 1;

-- name: GetSessionByCode :one
SELECT id, name, code, created_by, max_participants, status 
FROM sessions WHERE code = ? LIMIT 1;

-- name: CountUsersInSession :one
SELECT COUNT(*) FROM users WHERE session_id = ?;

-- name: CreateUser :execresult
INSERT INTO users (session_id, name) VALUES (?, ?);
