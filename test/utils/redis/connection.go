package redis

import (
	"sync"

	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/redis"
	redis2 "github.com/go-redis/redis/v8"
)

var (
	connectionOne sync.Once
	connection    *redis2.Client
)

func Connection() *redis2.Client {
	connectionOne.Do(func() {
		connection = redis.NewPool(NewConfig())
	})

	return connection
}
