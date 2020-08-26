package commands

import (
	"fmt"
	"net"

	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/controllers"
	"github.com/spf13/cobra"
)

func NewWhitelist(
	createHandler controllers.CreateWhiteListHandler,
	removeHandler controllers.RemoveWhiteListHandler,
) *cobra.Command {
	whitelistCmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Whitelist control",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cliWhitelist called")
		},
	}

	whitelistCmd.AddCommand(
		newWhitelistAdd(createHandler),
		newWhitelistRemove(removeHandler),
	)

	return whitelistCmd
}

func newWhitelistAdd(createHandler controllers.CreateWhiteListHandler) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add IP network to whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			subnet, err := cmd.Flags().GetIPNet("subnet")
			if err != nil {
				return err
			}

			if err := createHandler.Execute(usecases.AddIPToWhiteList{Subnet: subnet}); err != nil {
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

func newWhitelistRemove(removeHandler controllers.RemoveWhiteListHandler) *cobra.Command {
	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove IP network from whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			subnet, err := cmd.Flags().GetIPNet("subnet")
			if err != nil {
				return err
			}

			if err := removeHandler.Execute(usecases.RemoveIPFromWhiteList{Subnet: subnet}); err != nil {
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
