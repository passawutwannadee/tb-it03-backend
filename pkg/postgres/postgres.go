package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/passawutwannadee/tb-it03/config"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func Connect(ctx context.Context, c config.Database) (*Postgres, error) {

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", c.Username, c.Password, c.Name, c.Host, c.Port)

	// Initialize PostgreSQL
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Pool: pool,
	}, nil
}
