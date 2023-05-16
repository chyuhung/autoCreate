package openstack

import "github.com/gophercloud/gophercloud/openstack/identity/v3/projects"

// GetProjects 函数返回所有启用的项目列表
func (os *OpenStack) GetProjects() ([]projects.Project, error) {
	// 声明一个布尔类型的变量
	var enabled = true

	// 配置 ListOpts
	listOpts := projects.ListOpts{
		Enabled: &enabled,
	}

	// 获取所有项目的所有页
	allPages, err := projects.List(os.Keystone, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	// 从所有页中提取所有项目
	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		return nil, err
	}
	return allProjects, nil
}
