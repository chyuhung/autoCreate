package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// GetInstance 函数返回指定 ID 的虚拟机
func (os *openStack) GetInstance(id string) (*servers.Server, error) {
	// 从 Nova 服务中获取指定 ID 的虚拟机
	vm, err := servers.Get(os.Nova, id).Extract()
	if err != nil {
		return nil, err
	}

	// 返回虚拟机
	return vm, nil
}

// GetFlavors 函数返回所有属性为 public 的 Flavor
func (os *openStack) GetFlavors() ([]flavors.Flavor, error) {
	listOpts := flavors.ListOpts{
		AccessType: flavors.PublicAccess,
	}

	allPages, err := flavors.ListDetail(os.Nova, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}
	return allFlavors, nil
}

func (os *openStack) GetHypervisors() ([]hypervisors.Hypervisor, error) {
	listsOpts := hypervisors.ListOpts{}
	allPages, err := hypervisors.List(os.Nova, listsOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allHypervisors, err := hypervisors.ExtractHypervisors(allPages)
	if err != nil {
		return nil, err
	}
	return allHypervisors, nil
}

func (os *openStack) GetHypervisorNames() ([]string, error) {
	hypervisors, err := os.GetHypervisors()
	if err != nil {
		return nil, err
	}
	var hypervisorNames []string
	for _, h := range hypervisors {
		hypervisorNames = append(hypervisorNames, h.HypervisorHostname)
	}
	return hypervisorNames, nil
}
