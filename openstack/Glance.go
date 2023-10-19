package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

// GetImages 函数返回指定项目的所有镜像
func (os *OpenStack) GetImages(projectID string) ([]images.Image, error) {
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

func (os *OpenStack) GetImageId(name string) (string, error) {
	// 获取所有镜像
	allPages, err := images.List(os.Glance, images.ListOpts{}).AllPages()
	if err != nil {
		return "", err
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return "", err
	}

	// 查找指定名称的镜像
	for _, image := range allImages {
		if image.Name == name {
			return image.ID, nil
		}
	}

	return "", fmt.Errorf("no image available with the name")
}
