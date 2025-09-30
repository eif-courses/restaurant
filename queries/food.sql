-- name: GetAllFood :many
select * from foods order by name;

-- name: InsertFood :one
INSERT INTO foods (name, price) VALUES ($1, $2)
RETURNING *;