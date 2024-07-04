package auth

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rkuprov/mbot/pkg/cfg"
)

type Auth struct {
	pg *pgxpool.Pool
}

type TestAuth struct {
	*Auth
	Pool *pgxpool.Pool
}

func New(pCfg cfg.Postgres) (*Auth, error) {
	psql, err := pgxpool.New(context.Background(), fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		pCfg.User,
		pCfg.Password,
		pCfg.Host,
		pCfg.Port,
		pCfg.DBName,
		pCfg.SSLMode,
		pCfg.PoolMaxConns))
	if err != nil {
		return nil, err
	}
	return &Auth{
		pg: psql,
	}, nil
}

func NewTestAuth() (*TestAuth, func(), error) {
	configs, err := cfg.Load()
	if err != nil {
		return nil, nil, err
	}
	client, err := New(configs.Postgres)
	if err != nil {
		return nil, nil, err
	}
	return &TestAuth{
			Auth: client,
			Pool: client.pg,
		},
		func() {
			cleanup(client.pg)
		}, nil
}

func cleanup(psql *pgxpool.Pool) {
	ctx := context.Background()
	_, _ = psql.Exec(ctx, "TRUNCATE TABLE customers CASCADE")
	_, _ = psql.Exec(ctx, "TRUNCATE TABLE users CASCADE")
	_, _ = psql.Exec(ctx, "TRUNCATE TABLE session CASCADE")
	psql.Close()
}
