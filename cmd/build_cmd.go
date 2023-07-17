package cmd

import (
	"autoCreate/pkg/tools"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var input string // 实例信息源文件

func init() {
	buildCmd.Flags().StringVar(&ef.instanceFile, "output", "vm_info_format.txt", "File to store formatted VM information")
	buildCmd.Flags().StringVar(&input, "input", "", "File to store VM information")
	rootCmd.AddCommand(buildCmd)
}

// generate命令从一个csv文件中读取vm基本信息，格式化后输出到指定的csv文件中，通过修改该指定文件来约束vm的创建
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate formatted VM information from input file, and output to file, in CSV",
	Run: func(cmd *cobra.Command, args []string) {
		if input == "" {
			log.Fatalln("You must provide input file")
		}
		log.Info("inputFile:", input)
		log.Info("instanceFile:", ef.instanceFile)

		// 从文件中读取vm info
		log.Infoln("Reading information from file ...")
		vms, err := readCsvFile(input)
		if err != nil {
			log.Errorln(err)
		}
		var formattedVMArray []string
		// 格式化vm info,并写入CSV文件
		for _, v := range vms {
			line, err := GenerateVM(v)
			if err != nil {
				log.Errorln(err)
			}
			formattedVMArray = append(formattedVMArray, string(line))
		}
		log.Infoln("Writing information to file ...")
		tools.WriteToEnvFile(formattedVMArray, ef.instanceFile)

	}}

func GenerateVM(v vmInfo) ([]byte, error) {
	// 从文件中读取最佳匹配的image name
	imageNames, err := tools.ReadFromEnvFile(ef.imageFile)
	if err != nil {
		log.Warnln("no valid image name found.", err)
	}
	imageName, err := tools.FuzzyMatch(v.OsName, imageNames)
	if err != nil {
		log.Warnln("no matching image name found.", err)
	}
	// 从vm信息中构造flavor
	flavorName := v.Cpu + "C" + v.Mem + "G" + v.SysVolSize + "G"
	// 从vm信息中获取网络信息
	// 主要ip
	mainNet := v.VlanName + "," + v.Ipaddr
	// 其他ip
	var networks []string
	for k, v := range v.Networks {
		networks = append(networks, k+","+v)
	}
	// 获取hypervisor name
	hypervisors, err := tools.ReadFromEnvFile(ef.hypervisorFile)
	if err != nil {
		log.Warnln("no valid hypervisor name found.", err)
	}
	num, err := tools.GetRandIndex(hypervisors)
	if err != nil {
		log.Warnln("no available hypervisor name found.", err)
	}
	hypervisorName := hypervisors[num]
	// 获取卷类型名称
	volumeTypes, err := tools.ReadFromEnvFile(ef.volumeTypeFile)
	if err != nil {
		log.Warnln("no valid volume type found,", err)
	}
	randNum, err := tools.GetRandIndex(volumeTypes)
	if err != nil {
		log.Warnln("no available volume type name found.", err)
	}
	volumeType := "None"
	if randNum > -1 {
		volumeType = volumeTypes[randNum]
	}
	// CSV文件格式
	// ##镜像名称，规格名称，系统盘大小，数据盘大小，虚拟机名称，VLAN名称，IP地址，宿主机，卷类型
	return []byte(imageName + "," + flavorName + "," + v.SysVolSize + "," + v.DataVolSize + "," + v.Name + "," + mainNet + "," + hypervisorName + "," + volumeType), nil
}

// func GenerateCreateOpts(v vmInfo) (*bootfromvolume.CreateOptsExt, error) {
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
// 	return &bootfromvolume.CreateOptsExt{createOpts, blockVolumeArray}, nil
// }
