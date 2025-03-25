package postgresclient

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresClient(ctx context.Context, dsn string) (c *pgxpool.Pool, err error) {

	return pgxpool.New(ctx, dsn)
}
