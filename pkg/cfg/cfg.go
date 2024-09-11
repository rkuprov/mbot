package cfg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var subscriptionToken string

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

const (
	mac     = "darwin"
	linux   = "linux"
	windows = "windows"
)

func InstallPath() string {
	switch runtime.GOOS {
	case mac:
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "mBot")
	case linux:
	case windows:
		return filepath.Join("C:", "Program Files", "mBot")
	default:
		return ""
	}
	return ""
}

func BinaryPath() string {
	switch runtime.GOOS {
	case mac:
		return filepath.Join("System", "Applications")
	case linux:
	case windows:
		return filepath.Join(InstallPath())
	default:
		return ""
	}
	return ""
}

func BinaryName() string {
	switch runtime.GOOS {
	case mac:
		return "mBot.app"
	case linux:
	case windows:
		return "mBot.exe"
	default:
		return ""
	}
	return ""
}

func UpdateURL() string {
	switch os.Getenv("MBOT_ENV") {
	case "local":
		return "http://localhost:8080"
	case "testing":
		return "http://localhost:8080"
	default:
		return ""
	}
}

func SubscriptionToken() string {
	return subscriptionToken
}
