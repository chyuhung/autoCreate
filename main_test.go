package main

// import (
// 	"autoCreate/pkg/tools"
// 	"sort"
// 	"testing"
// 	"time"

// 	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
// 	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
// )

// var myClient = env.GetClient()

// var hypervisor1 = hypervisors.Hypervisor{
// 	CPUInfo:            hypervisors.CPUInfo{Arch: "x86_64"},
// 	CurrentWorkload:    1,
// 	Status:             "enabled",
// 	State:              "up",
// 	DiskAvailableLeast: 100,
// 	HostIP:             "10.0.0.1",
// 	FreeDiskGB:         50,
// 	FreeRamMB:          1024,
// 	HypervisorHostname: "hypervisor1",
// 	HypervisorType:     "QEMU",
// 	HypervisorVersion:  1005000,
// 	ID:                 "1",
// 	LocalGB:            100,
// 	LocalGBUsed:        50,
// 	MemoryMB:           2048,
// 	MemoryMBUsed:       1024,
// 	RunningVMs:         1,
// 	Service:            hypervisors.Service{Host: "compute", ID: "1"},
// 	VCPUs:              4,
// 	VCPUsUsed:          2,
// }

// var hypervisor2 = hypervisors.Hypervisor{
// 	CPUInfo:            hypervisors.CPUInfo{Arch: "x86_64"},
// 	CurrentWorkload:    2,
// 	Status:             "enabled",
// 	State:              "up",
// 	DiskAvailableLeast: 200,
// 	HostIP:             "10.0.0.2",
// 	FreeDiskGB:         100,
// 	FreeRamMB:          2048,
// 	HypervisorHostname: "hypervisor2",
// 	HypervisorType:     "QEMU",
// 	HypervisorVersion:  1006000,
// 	ID:                 "2",
// 	LocalGB:            200,
// 	LocalGBUsed:        100,
// 	MemoryMB:           4096,
// 	MemoryMBUsed:       2048,
// 	RunningVMs:         2,
// 	Service:            hypervisors.Service{Host: "compute", ID: "2"},
// 	VCPUs:              8,
// 	VCPUsUsed:          4,
// }

// var hypervisor3 = hypervisors.Hypervisor{
// 	CPUInfo:            hypervisors.CPUInfo{Arch: "x86_64"},
// 	CurrentWorkload:    3,
// 	Status:             "disabled",
// 	State:              "down",
// 	DiskAvailableLeast: 300,
// 	HostIP:             "10.0.0.3",
// 	FreeDiskGB:         150,
// 	FreeRamMB:          4096,
// 	HypervisorHostname: "hypervisor3",
// 	HypervisorType:     "QEMU",
// 	HypervisorVersion:  1007000,
// 	ID:                 "3",
// 	LocalGB:            300,
// 	LocalGBUsed:        150,
// 	MemoryMB:           8192,
// 	MemoryMBUsed:       4096,
// 	RunningVMs:         3,
// 	Service:            hypervisors.Service{Host: "compute", ID: "3"},
// 	VCPUs:              16,
// 	VCPUsUsed:          8,
// }

// var hypervisorsData = env.SortByFreeRamMB{hypervisor1, hypervisor2, hypervisor3, hypervisor1, hypervisor3}

// // 测试真实获取所有宿主机
// func TestGetHypervisors(t *testing.T) {
// 	list := myClient.GetHypervisors()
// 	for _, Hypervisor := range list {
// 		t.Log("CPUInfo:", Hypervisor.CPUInfo)
// 		t.Log("CurrentWorkload:", Hypervisor.CurrentWorkload)
// 		t.Log("Status:", Hypervisor.Status)
// 		t.Log("State:", Hypervisor.State)
// 		t.Log("DiskAvailableLeast:", Hypervisor.DiskAvailableLeast)
// 		t.Log("HostIP:", Hypervisor.HostIP)
// 		t.Log("FreeRamMB:", Hypervisor.FreeRamMB)
// 		t.Log("HypervisorHostname:", Hypervisor.HypervisorHostname)
// 		t.Log("HypervisorType:", Hypervisor.HypervisorType)
// 		t.Log("LocalGBUsed:", Hypervisor.LocalGBUsed)
// 		t.Log("MemoryMB:", Hypervisor.MemoryMB)
// 		t.Log("MemoryMBUsed:", Hypervisor.MemoryMBUsed)
// 		t.Log("RunningVMs:", Hypervisor.RunningVMs)
// 		t.Log("Service:", Hypervisor.Service)
// 		t.Log("Servers:", Hypervisor.Servers)
// 		t.Log("VCPUs:", Hypervisor.VCPUs)
// 		t.Log("VCPUsUsed:", Hypervisor.VCPUsUsed)
// 	}
// }

// // 测试显示所有instance
// func TestListInstance(t *testing.T) {
// 	var result []servers.Server
// 	result = myClient.ListInstances()
// 	for i, srv := range result {
// 		t.Log("------------\n", "this is instance ", i)
// 		t.Logf("ID: %s", srv.ID)
// 		t.Logf("TenantID: %s", srv.TenantID)
// 		t.Logf("UserID: %s", srv.UserID)
// 		t.Logf("Name: %s", srv.Name)
// 		t.Logf("Updated: %s", srv.Updated.Format(time.RFC3339))
// 		t.Logf("Created: %s", srv.Created.Format(time.RFC3339))
// 		t.Logf("HostID: %s", srv.HostID)
// 		t.Logf("Status: %s", srv.Status)
// 		t.Logf("Progress: %d", srv.Progress)
// 		t.Logf("AccessIPv4: %s", srv.AccessIPv4)
// 		t.Logf("AccessIPv6: %s", srv.AccessIPv6)
// 		t.Logf("Image: %v", srv.Image)
// 		t.Logf("Flavor: %v", srv.Flavor)
// 		t.Logf("Addresses: %v", srv.Addresses)
// 		t.Logf("Metadata: %v", srv.Metadata)
// 		t.Logf("Links: %v", srv.Links)
// 		t.Logf("KeyName: %s", srv.KeyName)
// 		t.Logf("AdminPass: %s", srv.AdminPass)
// 		t.Logf("SecurityGroups: %v", srv.SecurityGroups)
// 		t.Logf("AttachedVolumes: %v", srv.AttachedVolumes)
// 		t.Logf("Fault: %v", srv.Fault)
// 		t.Logf("Tags: %v", srv.Tags)
// 		t.Logf("ServerGroups: %v", srv.ServerGroups)
// 	}

// }

// // 创建测试案例，测试模拟获取宿主机
// func TestMoniGetHypervisors(t *testing.T) {

// 	t.Log(hypervisorsData)

// 	sort.Slice(hypervisorsData, func(i, j int) bool {
// 		return hypervisorsData.Less(i, j)
// 	})

// 	t.Log(hypervisorsData)
// }

// // 文件路径
// var hypervisorFilepath = "env_available_host_list.txt"
// var imageFilepath = "env_image_list.txt"
// var volumeTypeFilepath = "env_cinder_volume_type.txt"
// var instanceFilepath = "vm_info_format.txt"

// func TestWriteInstanceFormatFile(t *testing.T) {
// 	for i := 1; i < 4; i++ {
// 		var bestHypervisors env.SortByFreeRamMB
// 		//读取所有宿主机
// 		hypervisorList := file.ReadEnvFile(hypervisorFilepath)
// 		//fmt.Println(hypervisorList)
// 		if len(hypervisorList) == 0 {
// 			panic("No valid host available!")
// 		}
// 		bestHypervisors = env.GetBestHypervisors(hypervisorsData)
// 		//t.Logf("hypervisorsData:\n %v", hypervisorsData)
// 		sort.Sort(bestHypervisors)
// 		//t.Logf("bestHypervisors:\n %v", bestHypervisors)
// 		if len(bestHypervisors) != 0 {
// 			// 第几台宿主机，则取可用宿主机的第几个
// 			t.Logf("instance %d : hypervisor %s,Free mem:%d.", i, bestHypervisors[i%len(bestHypervisors)].HypervisorHostname+",", bestHypervisors[i%len(bestHypervisors)].FreeRamMB)
// 			//hypervisor = bestHypervisors[0].HypervisorHostname + ","
// 		} else { // 如果没有满足条件的宿主机，则从填写的全部宿主机中任选
// 			t.Log("没有满足要求的宿主机")
// 			t.Logf("instance %d : hypervisor %s", i, hypervisorList[tools.GetRandIndex(hypervisorList)]+",")
// 		}
// 	}
// }
