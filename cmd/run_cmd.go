package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(run_Cmd)
}

var run_Cmd = &cobra.Command{
	Use:   "run",
	Short: "Start to create instance based on the specified content file",
	Run: func(cmd *cobra.Command, args []string) {
	}}

// func createVM(v vmInfo) error {
// 	myOpenStack, err := openstack.NewOpenStack(conf)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// image id
// 	imageID, err := myOpenStack.GetImageIDByName(v.OsName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// flavor id
// 	flavorName := v.Cpu + "C" + v.Mem + "G" + v.SysVolSize + "G"
// 	flavorID, err := myOpenStack.GetFlavorIDByName(flavorName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// networks
// 	// main ip
// 	myNetworks := []servers.Network{}
// 	networkID, err := myOpenStack.GetNetworkIDByName(v.VlanName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	myNetworks = append(myNetworks, servers.Network{UUID: networkID, FixedIP: v.Ipaddr})
// 	// other ip
// 	for vlan, ip := range v.Networks {
// 		id, err := myOpenStack.GetNetworkIDByName(vlan)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		myNetworks = append(myNetworks, servers.Network{UUID: id, FixedIP: ip})
// 	}
// 	createOpts := &servers.CreateOpts{
// 		Name:             v.Name,
// 		ImageRef:         imageID,
// 		FlavorRef:        flavorID,
// 		AvailabilityZone: "Nova",
// 		Networks:         myNetworks,
// 	}
// 	// create volume
// 	blockVolumeArray := []bootfromvolume.BlockDevice{}
// }
