package cmd

import (
	"os"

	"autoCreate/openstack"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var conf openstack.OpenStackConfig

const (
	DefaultRegion = "RegionOne"
)

var rootCmd = &cobra.Command{
	Use:   "autoCreate",
	Short: "A simplified command-line utility for quickly and automatically generating OpenStack instances.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Fatalln("no valid flag")
		}
		return nil
	},
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})

	conf.Username = os.Getenv("OS_USERNAME")
	conf.Password = os.Getenv("OS_PASSWORD")
	conf.ProjectName = os.Getenv("OS_PROJECT_NAME")
	conf.DomainName = os.Getenv("OS_USER_DOMAIN_NAME")

	// Use the value of OS_REGION_NAME environment variable if it's set, otherwise use DefaultRegion
	region := os.Getenv("OS_REGION_NAME")
	if region == "" {
		region = DefaultRegion
	}
	conf.Region = region
	conf.AuthURL = os.Getenv("OS_AUTH_URL")
	if conf.Username == "" || conf.Password == "" || conf.ProjectName == "" || conf.DomainName == "" || conf.AuthURL == "" {
		log.Errorln("You must provide the username, password, project name, user domain name, and auth URL by source environment variables")
		os.Exit(1)
	}

	// 默认的存储文件信息
	ef.imageFile = "env_image_list.txt"
	ef.hypervisorFile = "env_available_host_list.txt"
	ef.volumeTypeFile = "env_cinder_volume_type.txt"
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err)
	}
}
