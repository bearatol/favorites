package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Pharmacies interface {
	SetPharmacy(ctx context.Context, user, pharmacy int64) error
	DeletePharmacy(ctx context.Context, user, pharmacy int64) error
}

type Products interface {
	SetProduct(ctx context.Context, user, product int64) error
	DeleteProduct(ctx context.Context, user, product int64) error
}

type Repository struct {
	Pharmacies
	Products
}

func NewRepository(db *pgxpool.Pool) (repo *Repository) {
	return &Repository{
		Pharmacies: NewPharmaciesPostgres(db),
		Products:   NewProductsPostgres(db),
	}
}
