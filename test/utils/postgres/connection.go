package postgres

import (
	"database/sql"
	"sync"

	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres"
)

var (
	connectionOne sync.Once
	connection    *sql.DB
)

func Connection() *sql.DB {
	connectionOne.Do(func() {
		connection = postgres.NewConnection(NewConfig())
	})

	return connection
}
