package productusecase

import "github.com/booscaaa/go-signoz-otl/core/domain"

type usecase struct {
	repository domain.ProductRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
