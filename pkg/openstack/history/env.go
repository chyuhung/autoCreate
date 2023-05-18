package openstack

// import (
// 	"sync"
// )

// // EnvFile 包含存储 OpenStack 创建实例可选参数的文件名
// type EnvFile struct {
// 	ImageFilename      string // 存储可用的 image name
// 	VolumeTypeFilename string // 存储可用的 volume type name
// 	HypervisorFilename string // 存储可用的 hypervisor name
// 	TempFilename       string // 存储待创建的 instance
// }

// var (
// 	once sync.Once
// 	env  *EnvFile
// )

// func init() {
// 	// 使用 sync.Once 确保只创建一个 EnvFile 实例
// 	once.Do(func() {
// 		env = &EnvFile{
// 			ImageFilename:      "env_image_list.txt",
// 			VolumeTypeFilename: "env_cinder_volume_type.txt",
// 			HypervisorFilename: "env_available_host_list.txt",
// 			TempFilename:       "vm_info_format.txt",
// 		}
// 	})
// }

// // Get 返回 EnvFile 实例的指针
// func GetEnv() *EnvFile {
// 	return env
// }

// // // 初始化环境变量
// // func (e *EnvFile) InitEnvFile(provider *gophercloud.ProviderClient) {
// // 	mu.Lock()
// // 	defer mu.Unlock()
// // 	// write image list to file
// // 	WriteEnvFile(openstack.Get().GetImageNames(), e.ImageFilename)
// // 	// write host list to file
// // 	hypervisorNames := make([]string, 0, len(myOpenstack.Get().GetHypervisors()))
// // 	for _, h := range myOpenstack.Get().GetHypervisors() {
// // 		hypervisorNames = append(hypervisorNames, h.HypervisorHostname)
// // 	}
// // 	WriteEnvFile(hypervisorNames, e.HypervisorFilename)
// // 	// write volume type list to file
// // 	WriteEnvFile(myOpenstack.Get().GetVolTypeNames(), e.VolumeTypeFilename)
// // }

// // // 写入缓存文件
// // func (e *EnvFile) WriteInstanceFormatFile() []string {
// // 	mu.Lock()
// // 	defer mu.Unlock()
// // 	instances := myOpenstack.ReadCsvFile()
// // 	// 固定格式，第一行数据便于阅读
// // 	firstLine := "#镜像," + "规格," + "系统盘Size," + "数据盘Size," + "实例名称," + "VLAN名称," + "IP地址," + "宿主机," + "卷类型,"
// // 	result := make([]string, 0, len(instances)+1)
// // 	result = append(result, firstLine)
// // 	for n, instance := range instances {
// // 		result = append(result, myOpenstack.Get().FormatInstanceInfo(instance, n))
// // 		common.Client.Get()
// // 	}
// // 	WriteEnvFile(result, e.TempFilename)
// // 	return result
// // }
