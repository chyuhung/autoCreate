package cmd

import (
	"autoCreate/pkg/openstack"
	"autoCreate/pkg/tools"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var ef envFile

func init() {
	initCmd.Flags().StringVar(&ef.hypervisorFile, "hypervisor", "env_available_host_list.txt", "File to store available hypervisor names")
	initCmd.Flags().StringVar(&ef.imageFile, "image", "env_image_list.txt", "File to store available image names")
	initCmd.Flags().StringVar(&ef.volumeTypeFile, "volume-type", "env_cinder_volume_type.txt", "File to store available volume type names")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Retrieve the names of the available hypervisors, image types, and volume types",
	Run: func(cmd *cobra.Command, args []string) {
		myOpenStack, err := openstack.NewOpenStack(conf)
		if err != nil {
			log.Fatalln(err)
		}
		var hypervisorNames []string
		var imageNames []string
		var volumeTypeNames []string

		// 获取 hypervisor name
		log.Info("Retrieving hypervisor names ...")
		hypervisorNames, err = myOpenStack.GetHypervisorNames()
		if err != nil {
			log.Errorln(err)
		}
		tools.WriteToEnvFile(tools.UniqueStrings(hypervisorNames), ef.hypervisorFile)
		log.Info("hypervisorFile:", ef.hypervisorFile)

		// 获取 image name
		log.Info("Retrieving image names ...")
		// 获取所有 project
		allProjects, err := myOpenStack.GetProjects()
		if err != nil {
			log.Fatalf("failed to get all projects: %v\n", err)
		}
		for _, p := range allProjects {
			if ef.imageFile != "" {
				images, err := myOpenStack.GetImages(p.ID)
				if err != nil {
					log.Errorln(err)
				}
				for _, i := range images {
					imageNames = append(imageNames, i.Name)
				}
				tools.WriteToEnvFile(tools.UniqueStrings(imageNames), ef.imageFile)
			}
		}
		log.Info("imageFile:", ef.imageFile)

		// 获取 volume type name
		log.Info("Retrieving volume type names ...")
		volumeTypeNames, err = myOpenStack.Cinder.GetVolumeTypeNames()
		if err != nil {
			log.Errorln(err)
		}
		tools.WriteToEnvFile(tools.UniqueStrings(volumeTypeNames), ef.volumeTypeFile)
		log.Info("volumeTypeFile:", ef.volumeTypeFile)
	},
}
