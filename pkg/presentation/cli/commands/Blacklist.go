package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewBlackList() *cobra.Command {
	var blacklistCmd = &cobra.Command{
		Use:   "blacklist",
		Short: "Blacklist control",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("blacklist called")
		},
	}

	blacklistCmd.AddCommand(
		newBlackListAdd(),
		newBlackListRemove(),
	)

	return blacklistCmd
}

func newBlackListAdd() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add IP network to blacklist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}

	return addCmd
}

func newBlackListRemove() *cobra.Command {
	// cliBlacklistRemoveCmd represents the cliBlacklistRemove command
	var removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove IP network from blacklist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove called")
		},
	}

	return removeCmd
}
