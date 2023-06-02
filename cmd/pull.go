package cmd

import (
	"autoCreate/pkg/openstack"
	"autoCreate/pkg/tools"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var ef envFile

func init() {
	pullCmd.Flags().StringVar(&ef.hypervisorFile, "hypervisor", "", "pull the hypervisor name to the specified file")
	pullCmd.Flags().StringVar(&ef.imageFile, "image", "", "pull the image name to the specified file")
	pullCmd.Flags().StringVar(&ef.volumeTypeFile, "volume-type", "", "pull the volume type name to the specified file")
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the specified content to the specified file",
	Run: func(cmd *cobra.Command, args []string) {
		if ef.hypervisorFile == "" && ef.imageFile == "" && ef.volumeTypeFile == "" {
			log.Fatalln("no valid flag")
		}
		os, err := openstack.NewOpenStack(conf)
		if err != nil {
			log.Fatalln(err)
		}
		var hypervisorNames []string
		var imageNames []string
		var volumeTypeNames []string

		// 获取 hypervisor name
		if ef.hypervisorFile != "" {
			log.Info("hypervisorFile:", ef.hypervisorFile)

			hypervisorNames, err = os.GetHypervisorNames()
			if err != nil {
				log.Errorln(err)
			}
			tools.WriteToEnvFile(tools.UniqueStrings(hypervisorNames), ef.hypervisorFile)
		}
		if ef.imageFile != "" {
			// 获取 image name
			log.Info("imageFile:", ef.imageFile)
			// 获取所有 project
			allProjects, err := os.GetProjects()
			if err != nil {
				log.Fatalf("failed to get all projects: %v\n", err)
			}
			for _, p := range allProjects {
				if ef.imageFile != "" {
					images, err := os.GetImages(p.ID)
					if err != nil {
						log.Errorln(err)
					}
					for _, i := range images {
						imageNames = append(imageNames, i.Name)
					}
					tools.WriteToEnvFile(tools.UniqueStrings(imageNames), ef.imageFile)
				}
			}
		}
		// 获取 volume type name
		if ef.volumeTypeFile != "" {
			log.Info("volumeTypeFile:", ef.volumeTypeFile)
			volumeTypeNames, err = os.Cinder.GetVolumeTypeNames()
			if err != nil {
				log.Errorln(err)
			}
			tools.WriteToEnvFile(tools.UniqueStrings(volumeTypeNames), ef.volumeTypeFile)
		}
	},
}
