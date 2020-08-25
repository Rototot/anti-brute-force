package commands

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/controllers"
	"github.com/spf13/cobra"
	"net"
)

func NewBlackList(
	createHandler controllers.CreateBlackListHandler,
	removeHandler controllers.RemoveBlackListHandler,
) *cobra.Command {
	var blacklistCmd = &cobra.Command{
		Use:   "blacklist",
		Short: "Blacklist control",
	}

	blacklistCmd.AddCommand(
		newBlackListAdd(createHandler),
		newBlackListRemove(removeHandler),
	)

	return blacklistCmd
}

func newBlackListAdd(
	createHandler controllers.CreateBlackListHandler,
) *cobra.Command {
	var addCmd = &cobra.Command{
		Use:     "add",
		Short:   "Add IP network to blacklist",
		Example: `blacklist add "192.168.1/8"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			subnet, err := cmd.Flags().GetIPNet("subnet")
			if err != nil {
				return err
			}

			if err := createHandler.Execute(usecases.AddIpToBlacklist{Subnet: subnet}); err != nil {
				return err
			}

			return nil
		},
	}

	addCmd.Flags().IPNet("subnet", net.IPNet{}, "Example 192.168.1.1/8")
	err := addCmd.MarkFlagRequired("subnet")
	if err != nil {
		panic(err)
	}

	return addCmd
}

func newBlackListRemove(
	removeHandler controllers.RemoveBlackListHandler,
) *cobra.Command {
	// cliBlacklistRemoveCmd represents the cliBlacklistRemove command
	var removeCmd = &cobra.Command{
		Use:     "remove",
		Short:   "Remove IP network from blacklist",
		Example: `blacklist remove "192.168.1/8"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			subnet, err := cmd.Flags().GetIPNet("subnet")
			if err != nil {
				return err
			}

			if err := removeHandler.Execute(usecases.RemoveIpFromBlackList{Subnet: subnet}); err != nil {
				return err
			}

			return nil
		},
	}

	removeCmd.Flags().IPNet("subnet", net.IPNet{}, "Example 192.168.1.1/8")
	err := removeCmd.MarkFlagRequired("subnet")
	if err != nil {
		panic(err)
	}

	return removeCmd
}
