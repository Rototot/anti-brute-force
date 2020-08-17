package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewLimiter() *cobra.Command {
	var limiterCmd = &cobra.Command{
		Use:   "limiter",
		Short: "Limiter control",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("limiter called")
		},
	}

	limiterCmd.AddCommand(
		newRateLimiterReset(),
	)

	return limiterCmd
}

func newRateLimiterReset() *cobra.Command {
	var resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset limits for ip and login",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("resetCmd called")
		},
	}

	return resetCmd
}
