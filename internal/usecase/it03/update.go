package it03

import (
	"context"

	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
)

func (u *useCase) Update(ctx context.Context, model *postgresrepo.IT03UpdateParams) ([]postgresrepo.IT03UpdateRow, error) {
	lists, err := u.pgRepo.IT03Update(ctx, *model)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
