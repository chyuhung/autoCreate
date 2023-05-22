package cmd

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var input string // 实例信息源文件

func init() {
	generateCmd.Flags().StringVar(&ef.instanceFile, "output", "", "Generate the instance content to the specified file")
	generateCmd.Flags().StringVar(&input, "input", "", "Generate the instance content from the specified file")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the instance content to the specified file",
	Run: func(cmd *cobra.Command, args []string) {
		if ef.instanceFile == "" || input == "" {
			log.Fatalln("You must provide input file and output file")
		}
		log.Info("inputFile:", input)
		log.Info("instanceFile:", ef.instanceFile)
	}}

func generateCreateOpts(instanceInfoFromCSV) error {
	createOpts := &servers.CreateOptsExt{
		CreateOptsBuilder: &servers.CreateOpts{
			Name:      "my-instance",
			FlavorRef: "yyy",
		},
		BootFromVolume: &bootfromvolume.BlockDevice{
			VolumeID:            "aaa",
			DeleteOnTermination: true,
			BootIndex:           0,
		},
	}

	server, err := servers.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created server %s\n", server.ID)
	return nil
}
