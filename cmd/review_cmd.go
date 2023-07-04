package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reviewCmd)
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Check if all content of the instance is available",
	Run: func(cmd *cobra.Command, args []string) {
	}}
