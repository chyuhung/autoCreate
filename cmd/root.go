package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd 是主命令
var rootCmd = &cobra.Command{
	Use:   "autoCreate",
	Short: "A streamlined command line utility for expeditiously and automatically generating OpenStack instances.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// init 函数初始化日志记录器
func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// Execute 函数执行主命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
