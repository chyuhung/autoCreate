package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

// GetImages 函数返回指定项目的所有镜像
func (os *openStack) GetImages(projectID string) ([]images.Image, error) {
	// 配置 ListOpts
	listOpts := images.ListOpts{
		Owner: projectID,
	}

	// 获取所有镜像的所有页
	allPages, err := images.List(os.Glance, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	// 从所有页中提取所有镜像
	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return nil, err
	}
	return allImages, nil
}
