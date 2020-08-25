package fixtures

import (
	"github.com/Rototot/anti-brute-force/test/utils/postgres"
	"github.com/go-testfixtures/testfixtures/v3"
)

func Load(paths ...string) {
	if len(paths) == 0 {
		paths = append(paths, "test/fixtures")
	}

	db := postgres.Connection()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths(paths...),
	)

	if err != nil {
		panic(err)
	}

	err = fixtures.Load()
	if err != nil {
		panic(err)
	}
}
