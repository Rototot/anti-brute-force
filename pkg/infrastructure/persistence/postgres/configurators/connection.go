package configurators

import (
	"database/sql"
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/infrastructure/configurators"
)
import _ "github.com/jackc/pgx/stdlib"

func NewDatabase(conf configurators.PostgresConfig) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// setting
	db.SetMaxOpenConns(20)

	return db, nil
}
