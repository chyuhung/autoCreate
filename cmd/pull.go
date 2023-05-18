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
		// Add code here to get the specified data and write it to the env file
		_, err := openstack.NewOpenStack(conf)
		if err != nil {
			log.Fatalln(err)
		}
		// 获取所有 project
		//os.GetProjects()
		// 获取 image name

		// 将 image name 写入文件
		tools.WriteToEnvFile([]string{"this is a image", "this is a image too"}, imageFile)

		// 获取 hypervisor name
		// 将 hypervisor name 写入文件
		tools.WriteToEnvFile([]string{"this is a hypervisor", "this is a hypervisor too"}, hypervisorFile)

		// 获取 volume type name
		// 将 volume type name 写入文件
		tools.WriteToEnvFile([]string{"this is a volume type", "this is a volume type too"}, volumeTypeFile)

		log.Info("hypervisorFile:", hypervisorFile)
		log.Info("volumeTypeFile:", volumeTypeFile)
		log.Info("imageFile:", imageFile)
	},
}
