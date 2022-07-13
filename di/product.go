package di

import (
	"github.com/booscaaa/go-signoz-otl/adapter/http/productservice"
	"github.com/booscaaa/go-signoz-otl/adapter/postgres/productrepository"
	"github.com/booscaaa/go-signoz-otl/core/domain"
	"github.com/booscaaa/go-signoz-otl/core/usecase/productusecase"
	"github.com/jmoiron/sqlx"
)

func ConfigProductDI(conn *sqlx.DB) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUsecase := productusecase.New(productRepository)
	productService := productservice.New(productUsecase)

	return productService
}
