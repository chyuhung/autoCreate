package openstack

// import (
// 	"autoCreate/pkg/tools"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// 	"sync"

// 	"github.com/gophercloud/gophercloud"
// 	"github.com/gophercloud/gophercloud/openstack"
// 	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
// )

// // 客户端
// var openstackClient *client

// var myEnv = env.Get()

// // 筛选宿主机内存值
// var BESTFREEMEMMB = 1024 * 1 //1 GB
// // volume none类型
// const NONE = "None"

// var mu sync.Mutex

// var envIdentityEndpoint = os.Getenv("OS_AUTH_URL")
// var envUsername = os.Getenv("OS_USERNAME")
// var envPassword = os.Getenv("OS_PASSWORD")
// var envProjectName = os.Getenv("OS_PROJECT_NAME")
// var envDomainName = os.Getenv("OS_PROJECT_DOMAIN_NAME")

// type client struct {
// 	Provider *gophercloud.ProviderClient
// }

// func (c client) Get() *client {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if openstackClient == nil {
// 		openstackClient = newClient()
// 	}
// 	return openstackClient
// }

// func newClient() *client {
// 	scop := gophercloud.AuthScope{
// 		ProjectName: envProjectName,
// 		DomainName:  envDomainName,
// 	}
// 	opts := gophercloud.AuthOptions{
// 		IdentityEndpoint: envIdentityEndpoint,
// 		Username:         envUsername,
// 		Password:         envPassword,
// 		DomainName:       envDomainName,
// 		Scope:            &scop,
// 	}
// 	provider, err := openstack.AuthenticatedClient(opts)
// 	if err != nil {
// 		fmt.Println("Authentication failed, please check the environment variables.")
// 		return nil
// 	}
// 	return &client{Provider: provider}
// }

// func checkEnvValue(values []string) {
// 	for _, v := range values {
// 		if len(v) == 0 {
// 			fmt.Printf("环境变量获取失败，请首先加载环境变量！\n")
// 			os.Exit(1)
// 		}
// 	}
// }

// // FormatInstanceInfo
// func (c *client) FormatInstanceInfo(instanceInfo InstanceInfo, numOfInstances int) string {
// 	var image, flavor, sysSize, dataSize, hypervisor, volType string

// 	// image
// 	image = instanceInfo.OsName + ","
// 	// flavor
// 	flavor = instanceInfo.Cpu + "C" + instanceInfo.Mem + "G" + instanceInfo.SysVolSize + "G" + ","
// 	// volume
// 	// 多个data volume支持，使用 "+" 进行分割，例如50+50,表示两块50G磁盘
// 	sysSize = instanceInfo.SysVolSize + ","
// 	dataSize = instanceInfo.DataVolSize + ","
// 	// name
// 	//sysVolName := instanceInfo.Name + "" + "sysvol"
// 	//dataVolName := instanceInfo.Name + "" + "datavol"
// 	// network
// 	//instanceInfo.VlanName
// 	//instanceInfo.Ipaddr
// 	// other network，多个网络支持，使用 ":"和";" 进行分割，例如vlan 20:10.191.2.3;vlan 10:10.191.3.2
// 	//instanceInfo.Networks

// 	// hypervisor
// 	hypervisorList := file.ReadEnvFile(myEnv.hypervisorFilepath)
// 	//fmt.Println(hypervisorList)
// 	if len(hypervisorList) == 0 {
// 		panic("No valid host available!")
// 	}
// 	var allAvailableHypervisors, bestHypervisors SortByFreeRamMB
// 	//获取所有宿主机
// 	allHypervisors := openstackClient.GetHypervisors()
// 	for _, he := range hypervisorList {
// 		for _, h := range allHypervisors {
// 			if h.HypervisorHostname == he {
// 				allAvailableHypervisors = append(allAvailableHypervisors, h)
// 			}
// 		}
// 	}
// 	bestHypervisors = GetBestHypervisors(allAvailableHypervisors)
// 	sort.Sort(bestHypervisors)
// 	if len(bestHypervisors) != 0 {
// 		// 第几台宿主机，则取可用宿主机的第几个
// 		hypervisor = bestHypervisors[numOfInstances%len(bestHypervisors)].HypervisorHostname + ","
// 		//hypervisor = bestHypervisors[0].HypervisorHostname + ","
// 	} else { // 如果没有满足条件的宿主机，则从填写的全部宿主机中任选
// 		hypervisor = hypervisorList[tools.GetRandIndex(hypervisorList)] + ","
// 	}

// 	// read volume type
// 	volumeList := file.ReadEnvFile(myEnv.volumeTypeFilepath)
// 	if len(volumeList) == 0 {
// 		volType = NONE + ","
// 	} else {
// 		lastNum := tools.GetIpLastNum(instanceInfo.Ipaddr)
// 		index := lastNum % len(volumeList)
// 		volType = volumeList[index] + ","
// 	}

// 	//return image + flavor + sysSize + dataSize + instanceInfo.Name + "," + instanceInfo.VlanName + "," + instanceInfo.Ipaddr + "," + hypervisor + volType
// 	//构造字符串
// 	var builder strings.Builder
// 	builder.WriteString(image)
// 	builder.WriteString(flavor)
// 	builder.WriteString(sysSize)
// 	builder.WriteString(dataSize)
// 	builder.WriteString(instanceInfo.Name)
// 	builder.WriteString(",")
// 	builder.WriteString(instanceInfo.VlanName)
// 	builder.WriteString(",")
// 	builder.WriteString(instanceInfo.Ipaddr)
// 	builder.WriteString(",")
// 	builder.WriteString(hypervisor)
// 	builder.WriteString(volType)
// 	return builder.String()
// }

// // 为Hypervisor实现sort接口
// type SortByFreeRamMB []hypervisors.Hypervisor

// func (a SortByFreeRamMB) Len() int           { return len(a) }
// func (a SortByFreeRamMB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortByFreeRamMB) Less(i, j int) bool { return a[i].FreeRamMB < a[j].FreeRamMB }

// // 筛选free mem大于设置阈值的宿主机
// func GetBestHypervisors(hypervisorList []hypervisors.Hypervisor) []hypervisors.Hypervisor {
// 	var result []hypervisors.Hypervisor
// 	for _, h := range hypervisorList {
// 		// 如果大于阈值，加入结果数组中
// 		if h.FreeRamMB > BESTFREEMEMMB {
// 			result = append(result, h)
// 		}
// 	}
// 	return result
// }
