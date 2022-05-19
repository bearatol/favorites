package service

import (
	"context"

	"github.com/bearatol/favorites/internal/repository"
)

type PharmaciesService struct {
	repo repository.Pharmacies
}

func NewPharmaciesService(repo repository.Pharmacies) *PharmaciesService {
	return &PharmaciesService{repo: repo}
}

func (p *PharmaciesService) SetPharmacy(ctx context.Context, user, pharmacy int64) error {
	if err := p.repo.SetPharmacy(ctx, user, pharmacy); err != nil {
		return err
	}
	return nil
}

func (p *PharmaciesService) DeletePharmacy(ctx context.Context, user, pharmacy int64) error {
	if err := p.repo.DeletePharmacy(ctx, user, pharmacy); err != nil {
		return err
	}
	return nil
}
