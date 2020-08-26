package redis

import (
	"fmt"

	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	"github.com/go-redis/redis/v8"
)

// todo rename
func NewPool(conf configurators.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: "",
		DB:       conf.DB,
		PoolSize: 500,
	})
}
