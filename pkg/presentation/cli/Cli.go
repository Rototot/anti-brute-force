package cli

import (
	"fmt"
	"github.com/Rototot/anti-brute-force/pkg/presentation/cli/commands"
	"github.com/spf13/cobra"
)

func NewCLiCmd() *cobra.Command {

	var CliCmd = &cobra.Command{
		Use:   "cliCmd",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a cliCmd library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Context()
			fmt.Println("cliCmd called")
		},
	}

	CliCmd.AddCommand(
		commands.NewBlackList(),
		commands.NewWhitelist(),
		commands.NewLimiter(),
	)

	return CliCmd
}
