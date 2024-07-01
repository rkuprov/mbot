package store

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rkuprov/mbot/pkg/cfg"
)

type Store struct {
	pg *pgxpool.Pool
}

type TestStore struct {
	*Store
	Pool *pgxpool.Pool
}

func New(pCfg cfg.Postgres) (*Store, error) {
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
	return &Store{
		pg: psql,
	}, nil
}

func NewTestStore() (*TestStore, func(), error) {
	configs, err := cfg.Load()
	if err != nil {
		return nil, nil, err
	}
	client, err := New(configs.Postgres)
	if err != nil {
		return nil, nil, err
	}
	return &TestStore{
			Store: client,
			Pool:  client.pg,
		},
		func() {
			cleanup(client.pg)
		}, nil
}

func cleanup(psql *pgxpool.Pool) {
	_, _ = psql.Exec(context.Background(), "TRUNCATE TABLE customers CASCADE")
	psql.Close()
}
