package cmd

import (
	"autoCreate/pkg/openstack"

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
	myOpenStack, err := openstack.NewOpenStack(conf)
	if err != nil {
		log.Fatalln(err)
	}
	// image id
	imageID, err := myOpenStack.GetImageIDByName(v.OsName)
	if err != nil {
		log.Println(err)
	}
	// flavor id
	flavorName := v.Cpu + "C" + v.Mem + "G" + v.SysVolSize + "G"
	flavorID, err := myOpenStack.GetFlavorIDByName(flavorName)
	if err != nil {
		log.Println(err)
	}
	// networks
	myNetworks := []servers.Network{}
	networkID, err := myOpenStack.GetNetworkIDByName(v.VlanName)
	if err != nil {
		log.Println(err)
	}
	myNetworks = append(myNetworks, servers.Network{UUID: networkID, FixedIP: v.Ipaddr})
	for vlan, ip := range v.Networks {
		id, err := myOpenStack.GetNetworkIDByName(vlan)
		if err != nil {
			log.Println(err)
		}
		myNetworks = append(myNetworks, servers.Network{UUID: id, FixedIP: ip})
	}
	return &servers.CreateOpts{
		Name:             v.Name,
		ImageRef:         imageID,
		FlavorRef:        flavorID,
		AvailabilityZone: "Nova",
		Networks:         myNetworks,
	}, nil
}
