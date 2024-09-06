package database

import (
	"context"
	"go-test/internal/modules/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Dbpool *pgxpool.Pool

func Init() {
	var err error
	Dbpool, err = pgxpool.New(context.Background(), config.AppConfig.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	// Check connection
	err = Dbpool.Ping(context.Background())
	if err != nil {
		panic(err)
	}
}
