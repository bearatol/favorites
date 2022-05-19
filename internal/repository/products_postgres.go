package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductsObj struct {
	db *pgxpool.Pool
}

func NewProductsPostgres(db *pgxpool.Pool) *ProductsObj {
	return &ProductsObj{db: db}
}

func (p *ProductsObj) SetProduct(ctx context.Context, user, product int64) error {
	time := time.Now()
	_, err := p.db.Exec(ctx, "INSERT INTO products(user_id, product_id, created_at, updated_at) VALUES($1, $2, $3, $3) ON CONFLICT (user_id, product_id) DO UPDATE SET updated_at = $3, deleted_at = NULL", user, product, time)
	return err
}

func (p *ProductsObj) DeleteProduct(ctx context.Context, user, product int64) error {
	time := time.Now()
	_, err := p.db.Exec(ctx, "UPDATE products SET deleted_at = $1 WHERE user_id = $2 AND product_id = $3", time, user, product)
	return err
}
