package configurators

import "github.com/spf13/viper"

const (
	defaultPostgresPort     = 5432
	defaultPostgresHost     = "postgres"
	defaultPostgresUser     = "app"
	defaultPostgresPassword = "app_pass"
	defaultPostgresDb       = "app"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     int
}

func NewPostgresConfig(v *viper.Viper) *PostgresConfig {
	conf := &PostgresConfig{
		Host:     v.GetString("APP_POSTGRES_HOST"),
		User:     v.GetString("APP_POSTGRES_USER"),
		Password: v.GetString("APP_POSTGRES_PASSWORD"),
		Dbname:   v.GetString("APP_POSTGRES_DB"),
		Port:     v.GetInt("APP_POSTGRES_PORT"),
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
		conf.Dbname = defaultPostgresDb
	}
	if conf.Port == 0 {
		conf.Port = defaultPostgresPort
	}

	return conf
}
