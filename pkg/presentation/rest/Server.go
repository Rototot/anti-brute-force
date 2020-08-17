package rest

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewServer() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "REST API command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
		},
	}

	return serverCmd
}
