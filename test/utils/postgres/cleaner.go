package postgres

import (
	"fmt"
	"sync"

	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var (
	cleanerOne sync.Once
	cleaner    dbcleaner.DbCleaner
)

var cleaningTables = []string{
	"public.whitelists",
	"public.blacklists",
}

func Clean() {
	cleanerOne.Do(func() {
		conf := NewConfig()

		cleaner = dbcleaner.New()

		eng := engine.NewPostgresEngine(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
		cleaner.SetEngine(eng)
	})

	cleaner.Clean(cleaningTables...)
}
