package postgres

import (
	"fmt"
	"sync"
)

var cleaningTables = []string{
	"public.whitelists",
	"public.blacklists",
}

func Clean() {
	var wg sync.WaitGroup
	conn := Connection()
	for _, table := range cleaningTables {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()

			_, err := conn.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", t))
			if err != nil {
				panic(err)
			}
		}(table)
	}

	wg.Wait()
}
