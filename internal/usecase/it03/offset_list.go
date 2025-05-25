package it03

import (
	"context"
	"fmt"

	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
)

type PaginatedList struct {
	Lists []postgresrepo.IT03ListRow
	// Total int64
}

func (u *useCase) OffsetList(ctx context.Context) (*PaginatedList, error) {
	lists, err := u.pgRepo.IT03List(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(lists)

	// total, err := u.pgRepo.IT03Count(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	if len(lists) == 0 {
		lists = []postgresrepo.IT03ListRow{}
	}

	return &PaginatedList{
		Lists: lists,
		// Total: total,
	}, nil
}
