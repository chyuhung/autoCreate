package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// attach volume to instance
func (os *OpenStack) AttachVolume(*servers.Server, string) error {
	return fmt.Errorf("")
}

// GetInstances 函数返回指定分页的虚拟机列表
func (os *OpenStack) GetInstances(pageSize int, marker string) ([]servers.Server, error) {
	var instances []servers.Server
	listOpts := servers.ListOpts{
		AllTenants: true,
		Limit:      pageSize,
		Marker:     marker, // 设置起始位置为空字符串
	}

	allPages, err := servers.List(os.Nova, listOpts).AllPages()
	if err != nil {
		return instances, nil
	}
	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		return instances, nil
	}
	instances = append(instances, allInstances...)
	return instances, nil
}

// GetInstance 函数返回指定 ID 的虚拟机
func (os *OpenStack) GetInstance(id string) (*servers.Server, error) {
	// 从 Nova 服务中获取指定 ID 的虚拟机
	vm, err := servers.Get(os.Nova, id).Extract()
	if err != nil {
		return nil, err
	}

	// 返回虚拟机
	return vm, nil
}

// GetFlavors 函数返回所有属性为 public 的 Flavor
func (os *OpenStack) GetFlavors() ([]flavors.Flavor, error) {
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

// GetHypervisors 函数返回所有 Hypervisor
func (os *OpenStack) GetHypervisors() ([]hypervisors.Hypervisor, error) {
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

// GetHypervisorNames 函数返回所有 Hypervisor names
func (os *OpenStack) GetHypervisorNames() ([]string, error) {
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

// CreateInstance 创建 instance
func (os *OpenStack) CreateInstance(createOpts servers.CreateOpts) (*servers.Server, error) {
	server, err := bootfromvolume.Create(os.Nova, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	return server, nil
}

// 通过flavor name 获取flavor id
func (os *OpenStack) GetFlavorIDByName(name string) (string, error) {
	listOpts := flavors.ListOpts{
		AccessType: flavors.PublicAccess,
	}

	allPages, err := flavors.ListDetail(os.Nova, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		panic(err)
	}

	for _, flavor := range allFlavors {
		if flavor.Name == name {
			return flavor.ID, nil
		}
	}
	return "", fmt.Errorf("no flavor available with the name")
}
