package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version 变量用于存储程序的版本号
var version = "1.0.0"

// versionCmd 是版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of the program.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
	},
}

// init 函数将版本命令添加到主命令中
func init() {
	rootCmd.AddCommand(versionCmd)
}
