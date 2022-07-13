package productrepository

import (
	"context"

	"github.com/booscaaa/go-signoz-otl/core/domain"
)

func (repository repository) Fetch(ctx context.Context) (*[]domain.Product, error) {
	products := []domain.Product{}

	err := repository.db.SelectContext(ctx, &products, "SELECT * FROM product;")

	if err != nil {
		return nil, err
	}

	return &products, nil
}
