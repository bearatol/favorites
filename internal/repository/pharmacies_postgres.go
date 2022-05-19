package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PharmaciesObj struct {
	db *pgxpool.Pool
}

func NewPharmaciesPostgres(db *pgxpool.Pool) *PharmaciesObj {
	return &PharmaciesObj{db: db}
}

func (p *PharmaciesObj) SetPharmacy(ctx context.Context, user, pharmacy int64) error {
	time := time.Now()
	_, err := p.db.Exec(ctx, "INSERT INTO pharmacies(user_id, pharmacy_id, created_at, updated_at) VALUES($1, $2, $3, $3) ON CONFLICT (user_id, pharmacy_id) DO UPDATE SET updated_at = $3, deleted_at = NULL", user, pharmacy, time)
	return err
}

func (p *PharmaciesObj) DeletePharmacy(ctx context.Context, user, pharmacy int64) error {
	time := time.Now()
	_, err := p.db.Exec(ctx, "UPDATE pharmacies SET deleted_at = $1 WHERE user_id = $2 AND pharmacy_id = $3", time, user, pharmacy)
	return err
}
