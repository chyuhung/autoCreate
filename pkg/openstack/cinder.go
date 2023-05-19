package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
)

// VolumeClient 接口定义了卷类型客户端的方法
type VolumeClient interface {
	List(opts interface{}) ([]interface{}, error)
}

// CinderV2Client 实现了 VolumeTypeClient 接口，用于访问 Cinder v2 的 API
type CinderV2Client struct {
	Client *gophercloud.ServiceClient
}

// List 实现了 VolumeTypeClient 接口中的 List 方法，用于获取所有卷类型
func (c *CinderV2Client) List(opts interface{}) ([]interface{}, error) {
	// 转换为 ListOptsV2 类型
	listOpts, ok := opts.(volumes.ListOpts)
	if !ok {
		return nil, fmt.Errorf("invalid list options")
	}

	// 调用 Cinder v2 API 获取卷类型
	allPages, err := volumes.List(c.Client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allVolumeTypes, err := volumes.ExtractVolumes(allPages)
	if err != nil {
		return nil, err
	}

	var volumeTypes []interface{}
	for _, vt := range allVolumeTypes {
		volumeTypes = append(volumeTypes, vt)
	}

	return volumeTypes, nil
}

// CinderV3Client 实现了 VolumeTypeClient 接口，用于访问 Cinder v3 的 API
type CinderV3Client struct {
	Client *gophercloud.ServiceClient
}

// List 实现了 VolumeTypeClient 接口中的 List 方法，用于获取所有卷类型
func (c *CinderV3Client) List(opts interface{}) ([]interface{}, error) {
	// 转换为 ListOpts 类型
	listOpts, ok := opts.(volumetypes.ListOpts)
	if !ok {
		return nil, fmt.Errorf("invalid list options")
	}

	// 调用 Cinder v3 API 获取卷类型
	allPages, err := volumetypes.List(c.Client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return nil, err
	}

	var volumeTypes []interface{}
	for _, vt := range allVolumeTypes {
		volumeTypes = append(volumeTypes, vt)
	}

	return volumeTypes, nil
}

// IsCinderV2 返回当前OpenStack环境的Cinder版本是否为v2
func (os *openStack) IsCinderV2() bool {
	return os.CinderV2 != nil
}

// IsCinderV3 返回当前OpenStack环境的Cinder版本是否为v3
func (os *openStack) IsCinderV3() bool {
	return os.CinderV3 != nil
}

// GetVolumeTypes 函数返回指定项目的所有卷类型
func (os *openStack) GetVolumeTypes(projectID string) ([]interface{}, error) {
	// 配置 ListOpts
	listOptsV2 := volumes.ListOpts{
		AllTenants: true,
	}

	listOptsV3 := volumetypes.ListOpts{}

	// 根据 Cinder 版本选择客户端
	var client VolumeClient
	if os.IsCinderV3() {
		client = &CinderV3Client{Client: os.CinderV2}
	} else if os.IsCinderV2() {
		client = &CinderV2Client{Client: os.CinderV3}
	} else {
		return nil, fmt.Errorf("no valid cinder client")
	}

	// 调用客户端的 List 方法获取所有卷类型
	allVolumeTypes, err := client.List(listOptsV3)
	if err != nil {
		allVolumeTypes, err = client.List(listOptsV2)
		if err != nil {
			return nil, err
		}
	}

	return allVolumeTypes, nil
}

func (os *openStack) GetVolumeTypeNames() ([]string, error) {
	allVolumeTypes, err := os.GetVolumeTypes("")
	if err != nil {
		return nil, err
	}
	var volumeTypeNames []string
	for _, vt := range allVolumeTypes {
		if vtn, ok := vt.(volumetypes.VolumeType); ok {
			volumeTypeNames = append(volumeTypeNames, vtn.Name)
		} else if vtn, ok := vt.(volumes.Volume); ok {
			volumeTypeNames = append(volumeTypeNames, vtn.Name)
		}
	}
	return volumeTypeNames, nil
}
