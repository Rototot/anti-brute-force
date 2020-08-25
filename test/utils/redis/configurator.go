package redis

import (
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	configurators2 "github.com/Rototot/anti-brute-force/test/utils/configurators"
)

func NewConfig() configurators.RedisConfig {
	return *configurators.NewRedisConfig(configurators2.NewViper())
}
