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

func Load() (*Cfg, error) {
	switch os.Getenv("MBOT_ENV") {
	case "local":
		path := "../../deployment/config.json"
		return getCFGFromPath(path)
	case "testing":
		// port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		// if err != nil {
		// 	return nil, err
		// }
		return &Cfg{
			Postgres: Postgres{
				Host:         os.Getenv("POSTGRES_HOST"),
				Port:         5432,
				User:         os.Getenv("POSTGRES_USER"),
				Password:     os.Getenv("POSTGRES_PASSWORD"),
				DBName:       os.Getenv("POSTGRES_DBNAME"),
				SSLMode:      "disable",
				PoolMaxConns: 10,
			},
		}, nil
	}

	return nil, fmt.Errorf("unknown environment")
}

func getCFGFromPath(path string) (*Cfg, error) {
	cfg := new(Cfg)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bts, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
