// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: block.sql

package sql_gen

import (
	"context"
)

const allBlocks = `-- name: AllBlocks :many
SELECT height, hash, timestamp FROM blocks
`

func (q *Queries) AllBlocks(ctx context.Context) ([]Block, error) {
	rows, err := q.db.Query(ctx, allBlocks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Block
	for rows.Next() {
		var i Block
		if err := rows.Scan(&i.Height, &i.Hash, &i.Timestamp); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteBlock = `-- name: DeleteBlock :one
DELETE FROM blocks WHERE height = $1 RETURNING height, hash, timestamp
`

func (q *Queries) DeleteBlock(ctx context.Context, height int64) (Block, error) {
	row := q.db.QueryRow(ctx, deleteBlock, height)
	var i Block
	err := row.Scan(&i.Height, &i.Hash, &i.Timestamp)
	return i, err
}

const findBlock = `-- name: FindBlock :one
SELECT height, hash, timestamp FROM blocks WHERE height = $1
`

func (q *Queries) FindBlock(ctx context.Context, height int64) (Block, error) {
	row := q.db.QueryRow(ctx, findBlock, height)
	var i Block
	err := row.Scan(&i.Height, &i.Hash, &i.Timestamp)
	return i, err
}

const getBlocks = `-- name: GetBlocks :many
SELECT height, hash, timestamp FROM blocks ORDER BY height DESC OFFSET $1 LIMIT $2
`

type GetBlocksParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetBlocks(ctx context.Context, arg GetBlocksParams) ([]Block, error) {
	rows, err := q.db.Query(ctx, getBlocks, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Block
	for rows.Next() {
		var i Block
		if err := rows.Scan(&i.Height, &i.Hash, &i.Timestamp); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestBlock = `-- name: GetLatestBlock :one
SELECT height, hash, timestamp FROM blocks ORDER BY height DESC LIMIT 1
`

func (q *Queries) GetLatestBlock(ctx context.Context) (Block, error) {
	row := q.db.QueryRow(ctx, getLatestBlock)
	var i Block
	err := row.Scan(&i.Height, &i.Hash, &i.Timestamp)
	return i, err
}

const insertBlock = `-- name: InsertBlock :one
INSERT INTO blocks (height, hash, timestamp) VALUES ($1, $2, $3) RETURNING height, hash, timestamp
`

type InsertBlockParams struct {
	Height    int64
	Hash      string
	Timestamp int64
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) (Block, error) {
	row := q.db.QueryRow(ctx, insertBlock, arg.Height, arg.Hash, arg.Timestamp)
	var i Block
	err := row.Scan(&i.Height, &i.Hash, &i.Timestamp)
	return i, err
}

const totalBlocks = `-- name: TotalBlocks :one
SELECT COUNT(*) FROM blocks
`

func (q *Queries) TotalBlocks(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, totalBlocks)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateBlock = `-- name: UpdateBlock :one
UPDATE blocks SET hash = $2, timestamp = $3 WHERE height = $1 RETURNING height, hash, timestamp
`

type UpdateBlockParams struct {
	Height    int64
	Hash      string
	Timestamp int64
}

func (q *Queries) UpdateBlock(ctx context.Context, arg UpdateBlockParams) (Block, error) {
	row := q.db.QueryRow(ctx, updateBlock, arg.Height, arg.Hash, arg.Timestamp)
	var i Block
	err := row.Scan(&i.Height, &i.Hash, &i.Timestamp)
	return i, err
}
