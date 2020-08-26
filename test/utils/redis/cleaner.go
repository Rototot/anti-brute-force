package redis

import "context"

func Clean() {
	conn := Connection()

	ctx := context.Background()

	err := conn.FlushAll(ctx).Err()
	if err != nil {
		panic(err)
	}
}
