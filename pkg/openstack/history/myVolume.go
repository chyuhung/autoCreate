package openstack

// import (
// 	"fmt"

// 	"github.com/gophercloud/gophercloud"
// 	"github.com/gophercloud/gophercloud/openstack"
// 	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
// )

// // create volume
// func (c *client) CreateVolume() {
// 	method := "CreateVolume"
// 	_, err := openstack.NewBlockStorageV2(c.Provider, gophercloud.EndpointOpts{
// 		Region: "RegionOne",
// 	})
// 	if err != nil {
// 		fmt.Printf("%s : %v", method, err)
// 		panic(err)
// 	}
// }

// // list volume type
// func (c *client) GetVolTypeNames() (volNames []string) {
// 	method := "getVolTypeNames"
// 	client, err := openstack.NewBlockStorageV2(c.Provider, gophercloud.EndpointOpts{
// 		Region: "RegionOne",
// 	})
// 	if err != nil {
// 		fmt.Printf("%s : %v", method, err)
// 		panic(err)
// 	}
// 	allPages, err := volumetypes.List(client, volumetypes.ListOpts{}).AllPages()
// 	if err != nil {
// 		panic(err)
// 	}
// 	volumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//没有数据的情况
// 	if len(volumeTypes) == 0 {
// 		volNames = append(volNames, NONE)
// 		return
// 	}
// 	for _, v := range volumeTypes {
// 		// 排除默认卷名称
// 		if v.Name == "__DEFAULT__" {
// 			volNames = append(volNames, NONE)
// 		} else {
// 			volNames = append(volNames, v.Name)
// 		}
// 		//fmt.Println(v.Name)
// 	}
// 	return
// }
