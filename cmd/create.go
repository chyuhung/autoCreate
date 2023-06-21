package cmd

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Start to create instance based on the specified content file",
	Run: func(cmd *cobra.Command, args []string) {
	}}

func createVM(opts servers.CreateOpts) (string, error) {
	return "", nil
}
