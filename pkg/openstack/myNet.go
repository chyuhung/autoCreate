package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/pagination"
)

// list network
func (c *client) GetNetworkNames() (networkNames []string) {
	method := "GetNetworkNames"
	client, err := openstack.NewNetworkV2(c.Provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	//fmt.Println(client)
	if err != nil {
		fmt.Printf("%s : %v", method, err)
		return
	}
	pager := networks.List(client, networks.ListOpts{})
	//fmt.Println(pager)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		extract, _ := networks.ExtractNetworks(page)
		for _, n := range extract {
			networkNames = append(networkNames, n.Name)
			//fmt.Println(n.Name)
		}
		return true, nil
	})
	return
}
