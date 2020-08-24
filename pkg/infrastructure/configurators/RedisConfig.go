package configurators

import "github.com/spf13/viper"

const (
	defaultRedisPort = 6379
	defaultRedisHost = ""
	defaultRedisDB   = 0
)

type RedisConfig struct {
	Host string
	Port int
	Db   int
}

func NewRedisConfig(v *viper.Viper) *RedisConfig {
	conf := &RedisConfig{
		Host: v.GetString("APP_REDIS_HOST"),
		Port: v.GetInt("APP_REDIS_PORT"),
		Db:   v.GetInt("APP_REDIS_DB"),
	}

	if conf.Port == 0 {
		conf.Port = defaultRedisPort
	}

	if conf.Db == 0 {
		conf.Db = defaultRedisDB
	}

	if conf.Host == "" {
		conf.Host = defaultRedisHost
	}

	return conf
}
