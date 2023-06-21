package cmd

import (
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

func GenerateCreateOpts(v vmInfo) (*servers.CreateOpts, error) {
	return &servers.CreateOpts{
		Name:      v.Name,
		ImageRef:  "",
		FlavorRef: "",
		//v.Cpu+"C"+v.Mem+"G"+v.SysVolSize+"G"
		AvailabilityZone: "Nova",
		Networks:         nil,
		AccessIPv4:       v.Ipaddr,
	}, nil
}
