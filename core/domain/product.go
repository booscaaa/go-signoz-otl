package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Product is entity of table product database column
type Product struct {
	ID   int32  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Fetch(*gin.Context)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Fetch(context.Context) (*[]Product, error)
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Fetch(context.Context) (*[]Product, error)
}
