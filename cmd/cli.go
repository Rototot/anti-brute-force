package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cliCmd represents the cliCmd command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cli app control",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli called")
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
