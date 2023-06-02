package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	volumesV2 "github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	volumesV3 "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
)

type CinderInterface interface {
	Create(opts interface{}) (interface{}, error)
	GetVolumeTypes() ([]interface{}, error)
	GetVolumeTypeNames() ([]string, error)
}

type Cinder struct {
	CinderServiceClient *gophercloud.ServiceClient
	Version             int // 添加版本信息
}

func (c *Cinder) Create(opts interface{}) (interface{}, error) {
	switch c.Version {
	case 2:
		if createOpts, ok := opts.(volumesV2.CreateOpts); ok {
			return volumesV2.Create(c.CinderServiceClient, createOpts).Extract()
		}
		return nil, fmt.Errorf("invalid options")
	case 3:
		if createOpts, ok := opts.(volumesV3.CreateOpts); ok {
			return volumesV3.Create(c.CinderServiceClient, createOpts).Extract()
		}
		return nil, fmt.Errorf("invalid options")
	default:
		return nil, fmt.Errorf("unsupported Cinder version")
	}
}

func (c *Cinder) GetVolumeTypes() ([]interface{}, error) {
	switch c.Version {
	case 2:
		listOpts := volumesV2.ListOpts{}
		// 调用 Cinder v2 API 获取卷类型
		allPages, err := volumesV2.List(c.CinderServiceClient, listOpts).AllPages()
		if err != nil {
			return nil, err
		}

		allVolumeTypes, err := volumesV2.ExtractVolumes(allPages)
		if err != nil {
			return nil, err
		}
		var result []interface{}
		result = append(result, allVolumeTypes)
		return result, nil
	case 3:
		listOpts := volumetypes.ListOpts{}
		allPages, err := volumetypes.List(c.CinderServiceClient, listOpts).AllPages()
		if err != nil {
			return nil, err
		}

		allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
		if err != nil {
			return nil, err
		}
		var result []interface{}
		result = append(result, allVolumeTypes)
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported Cinder version")
	}
}

// 获取所有 volume 类型的名称
func (c *Cinder) GetVolumeTypeNames() ([]string, error) {
	result := make(map[string]struct{}, 10)
	var volumeTypeNames []string
	switch c.Version {
	case 2:
		allVolumeTypes, err := c.GetVolumeTypes()
		if err != nil {
			return nil, err
		}
		for _, vt := range allVolumeTypes {
			if vtn, ok := vt.(volumetypes.VolumeType); ok {
				volumeTypeNames = append(volumeTypeNames, vtn.Name)
			} else if vtn, ok := vt.(volumesV2.Volume); ok {
				volumeTypeNames = append(volumeTypeNames, vtn.Name)
			}
		}
		for _, n := range volumeTypeNames {
			result[n] = struct{}{}
		}
		for k := range result {
			volumeTypeNames = append(volumeTypeNames, k)
		}
	}
	return volumeTypeNames, nil
}
