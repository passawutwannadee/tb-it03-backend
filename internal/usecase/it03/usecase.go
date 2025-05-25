package it03

import (
	"context"

	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
	"github.com/passawutwannadee/tb-it03/pkg/postgres"
)

type UseCase interface {
	OffsetList(ctx context.Context) (*PaginatedList, error)
	Update(ctx context.Context, model *postgresrepo.IT03UpdateParams) ([]postgresrepo.IT03UpdateRow, error)
}

type useCase struct {
	pg     *postgres.Postgres
	pgRepo *postgresrepo.Queries
}

func New(db *postgres.Postgres) UseCase {
	return &useCase{
		pg:     db,
		pgRepo: postgresrepo.New(db.Pool),
	}
}
