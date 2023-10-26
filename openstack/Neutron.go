package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

// Get network ids
func (os *OpenStack) BuildNetworks(networks map[string]string) ([]servers.Network, error) {
	var data []servers.Network
	// 遍历map
	for k, v := range networks {
		fmt.Printf("key: %s, value: %s\n", k, v)
		id, err := os.GetImageId(k)
		if err != nil {
			break
		}
		data = append(data, servers.Network{UUID: id, FixedIP: v})
	}
	return data, nil
}

// Get Networks
func (os *OpenStack) GetNetworks() ([]networks.Network, error) {
	// Create a ListOpts object
	listOpts := networks.ListOpts{}
	// Get all pages of the list
	allPages, err := networks.List(os.Neutron, listOpts).AllPages()
	if err != nil {
		// Return nil and the error if there is an error
		return nil, err
	}
	// Extract the networks from the allPages
	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		// Return nil and the error if there is an error
		return nil, err
	}
	// Return the allNetworks and nil
	return allNetworks, nil
}

// Get Network ID By Name
func (os *OpenStack) GetNetworkID(vlan string) (string, error) {
	// Get the networks
	networks, err := os.GetNetworks()
	if err != nil {
		// Return an empty string and the error if there is an error
		return "", err
	}
	// Iterate through the networks
	for _, n := range networks {
		// Check if the name matches
		if n.Name == vlan {
			// Return the ID and nil if there is a match
			return n.ID, nil
		}
	}
	// Return an empty string and nil if there is no match
	return "", nil
}
