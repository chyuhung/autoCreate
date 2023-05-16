package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Gets the specified data to the env file",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
