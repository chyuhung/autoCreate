package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	volumeTypeFile string
	hypervisorFile string
	imageFile      string
)

func init() {
	initCmd.Flags().StringVar(&hypervisorFile, "hypervisor", "", "init the file with hypervisor name")
	initCmd.Flags().StringVar(&imageFile, "image", "", "init the file with image name")
	initCmd.Flags().StringVar(&volumeTypeFile, "volumetype", "", "init the file with volume type name")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Gets the specified data and writes it to the env file",
	Run: func(cmd *cobra.Command, args []string) {
		// Add code here to get the specified data and write it to the env file
		fmt.Println("hypervisorFile:", hypervisorFile)
		fmt.Println("volumeTypeFile:", volumeTypeFile)
		fmt.Println("imageFile:", imageFile)
	},
}
