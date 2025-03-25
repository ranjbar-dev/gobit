-- name: AllBlocks :many
SELECT * FROM blocks;

-- name: GetBlocks :many
SELECT * FROM blocks ORDER BY height DESC OFFSET $1 LIMIT $2;

-- name: GetLatestBlock :one
SELECT * FROM blocks ORDER BY height DESC LIMIT 1;

-- name: FindBlock :one
SELECT * FROM blocks WHERE height = $1;

-- name: InsertBlock :one
INSERT INTO blocks (height, hash, timestamp) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateBlock :one
UPDATE blocks SET hash = $2, timestamp = $3 WHERE height = $1 RETURNING *;

-- name: DeleteBlock :one
DELETE FROM blocks WHERE height = $1 RETURNING *;

-- name: TotalBlocks :one
SELECT COUNT(*) FROM blocks;

    