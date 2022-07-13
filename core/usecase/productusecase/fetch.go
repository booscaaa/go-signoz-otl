package productusecase

import (
	"context"

	"github.com/booscaaa/go-signoz-otl/core/domain"
)

func (usecase usecase) Fetch(ctx context.Context) (*[]domain.Product, error) {
	products, err := usecase.repository.Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return products, err
}
