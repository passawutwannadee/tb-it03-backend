package it03

import (
	"context"

	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
)

type UseCase interface {
	OffsetList(ctx context.Context) (*PaginatedList, error)
	Update(ctx context.Context, model *postgresrepo.IT03UpdateParams) ([]postgresrepo.IT03UpdateRow, error)
}

type useCase struct {
	pgRepo *postgresrepo.Queries
}

func New(db postgresrepo.DBTX) UseCase {
	return &useCase{
		pgRepo: postgresrepo.New(db),
	}
}
