package postgres

import (
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
)

func NewConfig() configurators.PostgresConfig {
	return configurators.NewPostgresConfig()
}
