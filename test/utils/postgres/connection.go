package postgres

import (
	"database/sql"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/persistence/postgres"
	"sync"
)

var connectionOne sync.Once
var connection *sql.DB

func Connection() *sql.DB {
	connectionOne.Do(func() {
		connection = postgres.NewConnection(NewConfig())
	})

	return connection
}
