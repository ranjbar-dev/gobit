package repository

import (
	"context"

	"github.com/ranjbar-dev/gobit/config"
	"github.com/ranjbar-dev/gobit/internal/postgresclient"
	sql_gen "github.com/ranjbar-dev/gobit/sql/gen"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	sql            *sql_gen.Queries
	postgresClient *pgxpool.Pool
}

func (r *Repository) GetDatabaseConnection() *pgxpool.Pool {

	return r.postgresClient
}

func NewRepository() *Repository {

	client, err := postgresclient.NewPostgresClient(
		context.Background(),
		config.PostgresDsn(),
	)
	if err != nil {
		panic(err)
	}

	temp := Repository{
		postgresClient: client,
		sql:            sql_gen.New(client),
	}

	return &temp
}
