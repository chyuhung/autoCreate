package openstack

// import (
// 	"fmt"

// 	"github.com/gophercloud/gophercloud"
// 	"github.com/gophercloud/gophercloud/openstack"
// 	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
// )

// // list Hosts
// func (c *client) GetHypervisors() []hypervisors.Hypervisor {
// 	var hypervisorList []hypervisors.Hypervisor
// 	method := "GetHypervisors"
// 	client, err := openstack.NewComputeV2(c.Provider, gophercloud.EndpointOpts{Region: "RegionOne"})
// 	if err != nil {
// 		fmt.Printf("%s : %v", method, err)
// 		panic(err)
// 	}
// 	allPages, err := hypervisors.List(client, nil).AllPages()
// 	if err != nil {
// 		panic(err)
// 	}

// 	allHypervisors, err := hypervisors.ExtractHypervisors(allPages)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, h := range allHypervisors {
// 		//可用状态
// 		if h.Status == "enabled" && h.State == "up" {
// 			hypervisorList = append(hypervisorList, h)
// 		}
// 	}
// 	// fmt.Println("----------\n", hypervisorList)
// 	return hypervisorList
// }
