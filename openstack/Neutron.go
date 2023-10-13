package openstack

import "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"

// Get Networks
func (os *OpenStack) GetNetworks() ([]networks.Network, error) {
	listOpts := networks.ListOpts{}
	allPages, err := networks.List(os.Neutron, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return nil, err
	}
	return allNetworks, nil
}

// Get Network ID By Name
func (os *OpenStack) GetNetworkIDByName(name string) (string, error) {
	networks, err := os.GetNetworks()
	if err != nil {
		return "", err
	}
	for _, n := range networks {
		if n.Name == name {
			return n.ID, nil
		}
	}
	return "", nil
}
