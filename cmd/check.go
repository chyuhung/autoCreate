package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if all content of the instance is available",
	Run: func(cmd *cobra.Command, args []string) {
	}}
