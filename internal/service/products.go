package service

import (
	"context"

	"github.com/bearatol/favorites/internal/repository"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}

func (p *ProductsService) SetProduct(ctx context.Context, user, product int64) error {
	if err := p.repo.SetProduct(ctx, user, product); err != nil {
		return err
	}
	return nil
}

func (p *ProductsService) DeleteProduct(ctx context.Context, user, product int64) error {
	if err := p.repo.DeleteProduct(ctx, user, product); err != nil {
		return err
	}
	return nil
}
