package productrepository

import (
	"github.com/booscaaa/go-signoz-otl/core/domain"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

// New returns contract implementation of ProductRepository
func New(db *sqlx.DB) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
