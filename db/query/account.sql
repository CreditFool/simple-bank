-- name: CreateAccount :one
INSERT INTO account (
  owner,
  balance
) VALUES (
  $1, $2
) RETURNING *;
