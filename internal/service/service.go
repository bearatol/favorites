package service

import (
	"context"

	"github.com/bearatol/favorites/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Pharmacies interface {
	SetPharmacy(ctx context.Context, user, pharmacy int64) error
	DeletePharmacy(ctx context.Context, user, pharmacy int64) error
}

type Products interface {
	SetProduct(ctx context.Context, user, product int64) error
	DeleteProduct(ctx context.Context, user, product int64) error
}

type Service struct {
	Pharmacies
	Products
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Pharmacies: NewPharmaciesService(repo.Pharmacies),
		Products:   NewProductsService(repo.Products),
	}
}
