package openstack

// import (
// 	"fmt"
// 	"log"

// 	"github.com/gophercloud/gophercloud/openstack"
// 	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
// 	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
// )

// func main() {
// 	// 创建一个新的OpenStack客户端
// 	client, err := openstack.NewClient("")
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// 获取存储服务的API版本
// 	storageURL, err := client.()
// 	if err != nil {
// 		log.Fatalf("Failed to get storage URL: %v", err)
// 	}

// 	// 判断API版本
// 	switch storageURL.Path {
// 	case "/v2/":
// 		// 使用v2 API创建卷
// 		volume, err := createVolumeV2(client)
// 		if err != nil {
// 			log.Fatalf("Failed to create volume: %v", err)
// 		}
// 		fmt.Printf("Created volume: %s\n", volume.ID)
// 	case "/v3/":
// 		// 使用v3 API创建卷
// 		volume, err := createVolumeV3(client)
// 		if err != nil {
// 			log.Fatalf("Failed to create volume: %v", err)
// 		}
// 		fmt.Printf("Created volume: %s\n", volume.ID)
// 	default:
// 		log.Fatalf("Unsupported API version: %s", storageURL.Path)
// 	}
// }

// func createVolumeV2(client openstack.OpenStackClient) (string, error) {
// 	// 创建卷的配置
// 	volumeConfig := v2.CreateVolumeConfig{
// 		Name:       "test-volume",
// 		Size:       10,
// 		VolumeType: "nfs",
// 	}

// 	// 创建卷
// 	volume, err := v2.CreateVolume(client, volumeConfig)
// 	if err != nil {
// 		return "", err
// 	}

// 	return volume.ID, nil
// }

// func createVolumeV3(client openstack.OpenStackClient) (string, error) {
// 	// 创建卷的配置
// 	volumeConfig := v3.CreateVolumeConfig{
// 		Name:       "test-volume",
// 		Size:       10,
// 		VolumeType: "nfs",
// 	}

// 	// 创建卷
// 	volume, err := v3.CreateVolume(client, volumeConfig)
// 	if err != nil {
// 		return "", err
// 	}

// 	return volume.ID, nil
// }
