package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
)
import _ "github.com/jackc/pgx/stdlib"

func NewConnection(conf configurators.PostgresConfig) *sql.DB {

	dsn := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	db, err := sql.Open("pgx", dsn)
	// todo add repeater
	if err != nil {
		panic(err)
	}

	// setting
	db.SetMaxOpenConns(20)

	return db
}
