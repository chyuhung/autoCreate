package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	generateCmd.Flags().StringVar(&ef.instanceFile, "file", "", "Generate the instance information to the specified file")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the instance content to the specified file",
	Run: func(cmd *cobra.Command, args []string) {
		if ef.instanceFile == "" {
			log.Fatalln("no valid flag")
		}
		log.Info("instanceFile:", ef.instanceFile)
	}}
