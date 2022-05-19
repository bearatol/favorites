package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bearatol/favorites/pkg/config"
)

func NewPostgresDB(ctx context.Context) (db *pgxpool.Pool, err error) {
	connConfig, err := pgxpool.ParseConfig(config.Conf().Database.URL)
	if err != nil {
		return
	}

	db, err = pgxpool.ConnectConfig(ctx, connConfig)
	if err != nil {
		return
	}

	err = db.Ping(ctx)
	if err != nil {
		return
	}

	return
}
