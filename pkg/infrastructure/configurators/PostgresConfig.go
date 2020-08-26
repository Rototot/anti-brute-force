package configurators

import (
	"os"
	"strconv"
)

const (
	defaultPostgresPort           = 5432
	defaultPostgresHost           = "postgres"
	defaultPostgresUser           = "app"
	defaultPostgresPassword       = "app_pass"
	defaultPostgresDB             = "app"
	defaultPostgresMaxConnections = 20
)

type PostgresConfig struct {
	Host           string
	User           string
	Password       string
	Dbname         string
	Port           int
	MaxConnections int
}

func NewPostgresConfig() PostgresConfig {
	conf := PostgresConfig{
		Host:     os.Getenv("APP_POSTGRES_HOST"),
		User:     os.Getenv("APP_POSTGRES_USER"),
		Password: os.Getenv("APP_POSTGRES_PASSWORD"),
		Dbname:   os.Getenv("APP_POSTGRES_DB"),
	}
	var err error

	if os.Getenv("APP_POSTGRES_PORT") != "" {
		conf.Port, err = strconv.Atoi(os.Getenv("APP_POSTGRES_PORT"))
		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("APP_POSTGRES_MAX_CONNECTIONS") != "" {
		conf.MaxConnections, err = strconv.Atoi(os.Getenv("APP_POSTGRES_MAX_CONNECTIONS"))
		if err != nil {
			panic(err)
		}
	}

	if conf.Host == "" {
		conf.Host = defaultPostgresHost
	}
	if conf.User == "" {
		conf.User = defaultPostgresUser
	}
	if conf.Password == "" {
		conf.Password = defaultPostgresPassword
	}
	if conf.Dbname == "" {
		conf.Dbname = defaultPostgresDB
	}
	if conf.Port == 0 {
		conf.Port = defaultPostgresPort
	}

	if conf.MaxConnections == 0 {
		conf.MaxConnections = defaultPostgresMaxConnections
	}

	return conf
}
