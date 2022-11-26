package store

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"todo_app/config"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	// sqlx.Connectを使うと内部でpingする。
	db, err := sqlx.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	// openは実際に接続テストが行われない
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db.DB, "mysql")
	return xdb, func() { _ = db.Close() }, nil
}