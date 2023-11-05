package postgresql

import (
	"context"
	"fmt"

	"github.com/evgsrkn/go-ddd-example/user/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewConnectionPool(ctx context.Context, cfg *config.Cfg, log *zap.Logger) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Cannot create connection pool", zap.Error(err))
	}

	if _, err := conn.Acquire(ctx); err != nil {
		log.Fatal("Cannot create connection pool", zap.Error(err))
	}

	return conn
}
