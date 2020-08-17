package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewWhitelist() *cobra.Command {
	var whitelistCmd = &cobra.Command{
		Use:   "whitelist",
		Short: "Whitelist control",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cliWhitelist called")
		},
	}

	whitelistCmd.AddCommand(
		newWhitelistAdd(),
		newWhitelistRemove(),
	)

	return whitelistCmd
}

func newWhitelistAdd() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add IP network to whitelist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cliWhitelistAdd called")
		},
	}

	return addCmd
}

func newWhitelistRemove() *cobra.Command {
	var removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove IP network from whitelist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove called")
		},
	}

	return removeCmd
}
