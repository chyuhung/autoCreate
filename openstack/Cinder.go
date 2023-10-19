package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	volumetypesv1 "github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumetypes"
	volumetypesv3 "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
)

type CinderServiceClient struct {
	*CinderServiceClientV1
	*CinderServiceClientV2
	*CinderServiceClientV3
}

type CinderInterface interface {
	// Create(opts interface{}) (interface{}, error)
	GetVolumeTypes() ([]interface{}, error)
	GetVolumeTypeNames() ([]string, error)
}

type CinderServiceClientV1 struct {
	CinderServiceClient *gophercloud.ServiceClient
	Version             int // 添加版本信息
}

type CinderServiceClientV2 struct {
	CinderServiceClient *gophercloud.ServiceClient
	Version             int // 添加版本信息
}

type CinderServiceClientV3 struct {
	CinderServiceClient *gophercloud.ServiceClient
	Version             int // 添加版本信息
}

// 获取所有 volume types
func (c *CinderServiceClient) GetVolumeTypes() ([]interface{}, error) {
	if c.CinderServiceClientV1 != nil {
		allPages, err := volumetypesv1.List(c.CinderServiceClientV1.CinderServiceClient).AllPages()
		if err != nil {
			return nil, err
		}
		allVolumeTypes, err := volumetypesv1.ExtractVolumeTypes(allPages)
		if err != nil {
			return nil, err
		}
		var result []interface{}
		for _, vt := range allVolumeTypes {
			result = append(result, vt)
		}
		return result, nil
	}
	if c.CinderServiceClientV3 != nil {
		listOpts := volumetypesv3.ListOpts{}
		allPages, err := volumetypesv3.List(c.CinderServiceClientV3.CinderServiceClient, listOpts).AllPages()
		if err != nil {
			return nil, err
		}
		allVolumeTypes, err := volumetypesv3.ExtractVolumeTypes(allPages)
		if err != nil {
			return nil, err
		}
		var result []interface{}
		for _, vt := range allVolumeTypes {
			result = append(result, vt)
		}
		return result, nil
	}
	return nil, fmt.Errorf("unsupported Cinder version")
}

// 获取所有 volume 类型的名称
func (c *CinderServiceClient) GetVolumeTypeNames() ([]string, error) {
	var volumeTypeNames []string
	allVolumeTypes, err := c.GetVolumeTypes()
	if err != nil {
		return nil, err
	}
	//log.Println(allVolumeTypes...)
	if allVolumeTypes == nil {
		return nil, fmt.Errorf("failed to get volume types")
	}
	for _, vt := range allVolumeTypes {
		switch v := vt.(type) {
		case volumetypesv1.VolumeType:
			v, ok := vt.(volumetypesv1.VolumeType)
			if !ok {
				fmt.Printf("Error: unable to convert %v to volumesV2.Volume\n", vt)
				continue
			}
			volumeTypeNames = append(volumeTypeNames, v.Name)
		case volumetypesv3.VolumeType:
			v, ok := vt.(volumetypesv3.VolumeType)
			if !ok {
				fmt.Printf("Error: unable to convert %v to volumetypes.VolumeType\n", vt)
				continue
			}
			volumeTypeNames = append(volumeTypeNames, v.Name)
		default:
			fmt.Printf("Error: unknown volume type %v\n", vt)
		}
	}
	return volumeTypeNames, nil
}

func (os *OpenStack) GetVolumeTypeId(string) (string, error) {
	return "", fmt.Errorf("")
}

func (os *OpenStack) CreateVolumes(map[string]int, string) ([]string, error) {
	return []string{}, fmt.Errorf("")
}

// 创建volume
// func (c *Cinder) Create() ( error) {
// 	bootfromvolume.CreateOptsExt{}
// 	return volumesV2.Create(c.CinderServiceClient, opts).Extract()
// }

// type VolumeOpts struct {
// 	// The size of the volume, in GB
// 	Size int `json:"image_id,required:true"`
// 	// The volume name
// 	Name string `json:"image_id,required:true"`
// 	// The volume type
// 	VolumeType string `json:"image_id,omitempty"`
// 	// The image id of the volume
// 	ImageID string `json:"image_id,omitempty"`
// }

// // 创建volume
// func (c *Cinder) GetCreateOpts(opts VolumeOpts) (interface{}, error) {
// 	return &volumesV2.CreateOpts{
// 		Size:       opts.Size,
// 		Name:       opts.Name,
// 		ImageID:    opts.ImageID,
// 		VolumeType: opts.VolumeType,
// 	}, nil
// }

// 创建volume
// func (c *Cinder) Create(opts interface{}) (interface{}, error) {
// 	switch c.Version {
// 	case 2:
// 		if createOpts, ok := opts.(volumesV2.CreateOpts); ok {
// 			return volumesV2.Create(c.CinderServiceClient, createOpts).Extract()
// 		}
// 		return nil, fmt.Errorf("invalid options")
// 	case 3:
// 		if createOpts, ok := opts.(volumesV3.CreateOpts); ok {
// 			return volumesV3.Create(c.CinderServiceClient, createOpts).Extract()
// 		}
// 		return nil, fmt.Errorf("invalid options")
// 	default:
// 		return nil, fmt.Errorf("unsupported Cinder version")
// 	}
// }

// // 获取所有 volume types
// func (c *CinderServiceClientV2) GetVolumeTypes() ([]interface{}, error) {
// 	switch c.Version {
// 	case 2:
// 		listOpts := volumesV2.ListOpts{AllTenants: true}
// 		allPages, err := volumesV2.List(c.CinderServiceClient, listOpts).AllPages()
// 		if err != nil {
// 			return nil, err
// 		}
// 		allVolumeTypes, err := volumesV2.ExtractVolumes(allPages)
// 		if err != nil {
// 			return nil, err
// 		}
// 		var result []interface{}
// 		for _, vt := range allVolumeTypes {
// 			result = append(result, vt)
// 		}
// 		return result, nil
// 	case 3:
// 		listOpts := volumetypes.ListOpts{}
// 		allPages, err := volumetypes.List(c.CinderServiceClient, listOpts).AllPages()
// 		if err != nil {
// 			return nil, err
// 		}

// 		allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
// 		if err != nil {
// 			return nil, err
// 		}
// 		var result []interface{}
// 		for _, vt := range allVolumeTypes {
// 			result = append(result, vt)
// 		}
// 		return result, nil
// 	default:
// 		return nil, fmt.Errorf("unsupported Cinder version")
// 	}
// }

// // 获取所有 volume 类型的名称
// func (c *CinderServiceClientV2) GetVolumeTypeNames() ([]string, error) {
// 	var volumeTypeNames []string
// 	allVolumeTypes, err := c.GetVolumeTypes()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("allVolumeTypes:", allVolumeTypes)
// 	for _, vt := range allVolumeTypes {
// 		switch v := vt.(type) {
// 		case volumesV2.Volume:
// 			v, ok := vt.(volumesV2.Volume)
// 			if !ok {
// 				fmt.Printf("Error: unable to convert %v to volumesV2.Volume\n", vt)
// 				continue
// 			}
// 			volumeTypeNames = append(volumeTypeNames, v.VolumeType)
// 			fmt.Println("typeName:", v.VolumeType)
// 		case volumetypes.VolumeType:
// 			v, ok := vt.(volumetypes.VolumeType)
// 			if !ok {
// 				fmt.Printf("Error: unable to convert %v to volumetypes.VolumeType\n", vt)
// 				continue
// 			}
// 			volumeTypeNames = append(volumeTypeNames, v.Name)
// 			fmt.Println("typeName:", v.Name)
// 		default:
// 			fmt.Printf("Error: unknown volume type %v\n", vt)
// 		}
// 	}
// 	return volumeTypeNames, nil
// }
