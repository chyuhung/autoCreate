package cmd

import (
	"os"

	"autoCreate/pkg/openstack"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var conf openstack.OpenStackConfig

const (
	DefaultRegion         = "RegionOne"
	DefaultUsername       = "OS_USERNAME"
	DefaultPassword       = "OS_PASSWORD"
	DefaultProjectName    = "OS_PROJECT_NAME"
	DefaultUserDomainName = "OS_USER_DOMAIN_NAME"
	DefaultAuthURL        = "OS_AUTH_URL"
)

var rootCmd = &cobra.Command{
	Use:   "autoCreate",
	Short: "A streamlined command line utility for expeditiously and automatically generating OpenStack instances.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})

	rootCmd.PersistentFlags().StringVarP(&conf.Username, "username", "u", os.Getenv(DefaultUsername), "username")
	rootCmd.PersistentFlags().StringVarP(&conf.Password, "password", "p", os.Getenv(DefaultPassword), "password")
	rootCmd.PersistentFlags().StringVarP(&conf.ProjectName, "project name", "j", os.Getenv(DefaultProjectName), "project name")
	rootCmd.PersistentFlags().StringVarP(&conf.DomainName, "domain name", "d", os.Getenv(DefaultUserDomainName), "domain name")

	// Use the value of OS_REGION_NAME environment variable if it's set, otherwise use DefaultRegion
	var region string
	if region = os.Getenv("OS_REGION_NAME"); region == "" {
		region = DefaultRegion
	}
	rootCmd.PersistentFlags().StringVarP(&conf.Region, "region", "r", region, "region")
	rootCmd.PersistentFlags().StringVarP(&conf.AuthURL, "auth url", "a", os.Getenv(DefaultAuthURL), "auth url")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
