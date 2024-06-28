package cfg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Cfg struct {
	Postgres Postgres `json:"postgres"`
}

type Postgres struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DBName       string `json:"dbname"`
	SSLMode      string `json:"sslmode"`
	PoolMaxConns int    `json:"pool_max_conns"`
}

func (c *Cfg) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	setupEnv(c)

	return json.Unmarshal(bts, c)
}

func setupEnv(c *Cfg) {
	err := os.Setenv("GOOSE_DRIVER", "postgres")
	if err != nil {
		panic(err)
	}
	// GOOSE_DRIVER=DRIVER
	err = os.Setenv("GOOSE_DBSTRING", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DBName,
		c.Postgres.SSLMode,
		c.Postgres.PoolMaxConns,
	))
	if err != nil {
		panic(err)
	}
	// GOOSE_MIGRATION_DIR=MIGRATION_DIR
	err = os.Setenv("GOOSE_MIGRATION_DIR", "../../migrations")
}
