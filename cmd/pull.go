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
	pullCmd.Flags().StringVar(&hypervisorFile, "hypervisor", "", "pull the hypervisor name to the specified file")
	pullCmd.Flags().StringVar(&imageFile, "image", "", "pull the image name to the specified file")
	pullCmd.Flags().StringVar(&volumeTypeFile, "volumetype", "", "pull the volume type name to the specified file")
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the specified data to the specified file",
	Run: func(cmd *cobra.Command, args []string) {
		// Add code here to get the specified data and write it to the env file
		fmt.Println("hypervisorFile:", hypervisorFile)
		fmt.Println("volumeTypeFile:", volumeTypeFile)
		fmt.Println("imageFile:", imageFile)
	},
}
