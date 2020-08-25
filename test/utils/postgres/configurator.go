package postgres

import (
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
	configurators2 "github.com/Rototot/anti-brute-force/test/utils/configurators"
)

func NewConfig() configurators.PostgresConfig {
	return *configurators.NewPostgresConfig(configurators2.NewViper())
}
