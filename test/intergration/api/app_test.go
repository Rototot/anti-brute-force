//+build e2e

package api

import (
	"context"
	"github.com/Rototot/anti-brute-force/cmd"
)

const baseUrl = "http://127.0.0.1:80"

func StartApp() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	c := cmd.NewServerCmd()

	go c.ExecuteContext(ctx)

	return ctx, cancel
}
