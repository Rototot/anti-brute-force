package configurators

import (
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewPool(conf configurators.RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port)) },
	}
}
