package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/pagination"
)

// list flavor
func (c *client) GetFlavorNames() (flavorNames []string) {
	method := "GetFlavorNames"
	client, err := openstack.NewComputeV2(c.Provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	//fmt.Println(client)
	if err != nil {
		fmt.Printf("%s : %v", method, err)
		return
	}
	pager := flavors.ListDetail(client, flavors.ListOpts{})
	//fmt.Println(pager)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		extract, _ := flavors.ExtractFlavors(page)
		for _, f := range extract {
			flavorNames = append(flavorNames, f.Name)
			fmt.Println(f.Name)
		}
		return true, nil
	})
	return
}
