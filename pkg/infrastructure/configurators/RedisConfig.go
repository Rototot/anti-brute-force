package configurators

import (
	"os"
	"strconv"
)

const (
	defaultRedisPort = 6379
	defaultRedisHost = ""
	defaultRedisDB   = 0
)

type RedisConfig struct {
	Host string
	Port int
	DB   int
}

func NewRedisConfig() RedisConfig {
	var err error
	var conf RedisConfig

	conf.Host = os.Getenv("APP_REDIS_HOST")
	if os.Getenv("APP_REDIS_PORT") == "" {
		conf.Port, err = strconv.Atoi(os.Getenv("APP_REDIS_PORT"))
		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("APP_REDIS_DB") == "" {
		conf.DB, err = strconv.Atoi(os.Getenv("APP_REDIS_DB"))
		if err != nil {
			panic(err)
		}
	}

	if conf.Port == 0 {
		conf.Port = defaultRedisPort
	}

	if conf.DB == 0 {
		conf.DB = defaultRedisDB
	}

	if conf.Host == "" {
		conf.Host = defaultRedisHost
	}

	return conf
}
