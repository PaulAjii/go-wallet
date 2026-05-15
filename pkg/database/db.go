package database

import (
	"context"
	"log"

	"github.com/PaulAjii/go-wallet/pkg/config"
	"github.com/PaulAjii/go-wallet/pkg/sysmsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

var Pool *pgxpool.Pool

func Connect() {
	var err error

	config, _ := pgxpool.ParseConfig(config.ApplicationConfig.Database.URL)

	config.ConnConfig.StatementCacheCapacity = 0
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("%s: %v\n", sysmsg.CannotConnect, err)
	}

	if err = Pool.Ping(context.Background()); err != nil {
		log.Fatalf("%s: %v\n", sysmsg.CannotPing, err)
	}

	log.Println(sysmsg.ConnectionSuccessful)
}

func Close() {
	Pool.Close()
	log.Println(sysmsg.ConnectionClosed)
}
