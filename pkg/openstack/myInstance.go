package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// create instance
func (c *client) CreateInstance(name string) {

	fmt.Println("create instance..........")
	client, err := openstack.NewComputeV2(c.Provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	ss, err := servers.Create(client, servers.CreateOpts{
		Name:             name,
		FlavorRef:        "64",
		ImageRef:         "aba55c25-60b5-4097-882d-9625fbc8fc11",
		AvailabilityZone: "nova",
		Networks: []servers.Network{
			servers.Network{UUID: "79e9b0c2-7e49-49f9-a7e2-0e03dd2055b3"},
		},
		AdminPass: "root",
	}).Extract()

	if err != nil {
		fmt.Printf("Create : %v", err)
		return
	}
	fmt.Println(ss)
}

// list instance
func (c *client) ListInstances() (result []servers.Server) {
	method := "ListInstance"
	client, err := openstack.NewComputeV2(c.Provider, gophercloud.EndpointOpts{Region: "RegionOne"})
	if err != nil {
		fmt.Printf("%s : %v", method, err)
		panic(err)
	}
	listOpts := servers.ListOpts{
		AllTenants: true,
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	result, err = servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}
	return result
}
