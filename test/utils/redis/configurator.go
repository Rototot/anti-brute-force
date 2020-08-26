package redis

import (
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
)

func NewConfig() configurators.RedisConfig {
	return configurators.NewRedisConfig()
}
