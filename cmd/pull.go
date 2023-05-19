package cmd

import (
	"autoCreate/pkg/openstack"
	"autoCreate/pkg/tools"

	log "github.com/sirupsen/logrus"

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
		if hypervisorFile == "" || imageFile == "" || volumeTypeFile == "" {
			log.Fatalln("no valid flag")
		}

		log.Info("hypervisorFile:", hypervisorFile)
		log.Info("volumeTypeFile:", volumeTypeFile)
		log.Info("imageFile:", imageFile)

		os, err := openstack.NewOpenStack(conf)
		if err != nil {
			log.Fatalln(err)
		}
		// 获取所有 project
		allProjects, err := os.GetProjects()
		if err != nil {
			log.Fatalf("failed to get all projects: %v\n", err)
		}
		var hypervisorNames []string
		var imageNames []string
		var volumeTypeNames []string

		for _, p := range allProjects {
			// 获取 image name
			images, err := os.GetImages(p.ID)
			if err != nil {
				log.Errorln(err)
			}
			for _, i := range images {
				imageNames = append(imageNames, i.Name)
			}
		}
		// 获取 hypervisor name
		hypervisorNames, err = os.GetHypervisorNames()
		if err != nil {
			log.Errorln(err)
		}
		// 获取 volume type name
		volumeTypeNames, err = os.GetVolumeTypeNames()
		if err != nil {
			log.Errorln(err)
		}

		// 将所有数据写入文件
		tools.WriteToEnvFile(tools.UniqueStrings(imageNames), imageFile)
		tools.WriteToEnvFile(tools.UniqueStrings(hypervisorNames), hypervisorFile)
		tools.WriteToEnvFile(tools.UniqueStrings(volumeTypeNames), volumeTypeFile)
	},
}
