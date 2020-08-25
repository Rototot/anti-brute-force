package redis

import "context"

func Clean() {
	conn := Connection()

	var ctx = context.Background()

	err := conn.FlushAll(ctx).Err()
	if err != nil {
		panic(err)
	}
}
