package openstack

import (
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	log "github.com/sirupsen/logrus"
)

// 默认区域名称
const DefaultRegion = "RegionOne"

// OpenStack 结构体包含各个服务的客户端
type OpenStack struct {
	Nova     *gophercloud.ServiceClient
	Neutron  *gophercloud.ServiceClient
	Glance   *gophercloud.ServiceClient
	CinderV2 *gophercloud.ServiceClient
	CinderV3 *gophercloud.ServiceClient
	Keystone *gophercloud.ServiceClient
}

// NewOpenStack 函数初始化 OpenStack 结构体
func NewOpenStack() (*OpenStack, error) {
	// 从环境变量中获取认证选项
	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}
	// 允许重新认证
	authOpts.AllowReauth = true

	// 认证，获取 provider
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return nil, err
	}

	// 获取区域名称
	region := os.Getenv("OS_REGION_NAME")
	if region == "" {
		region = DefaultRegion
	}

	// 初始化 OpenStack 结构体
	os := OpenStack{}
	os.Nova, err = getComputeClient(provider, region)
	if err != nil {
		return nil, err
	}
	os.Neutron, err = getNetworkClient(provider, region)
	if err != nil {
		return nil, err
	}
	os.Glance, err = getImageClient(provider, region)
	if err != nil {
		return nil, err
	}
	os.CinderV2, err = getBlockStorageV2Client(provider, region)
	if err != nil {
		return nil, err
	}
	os.CinderV3, err = getBlockStorageV3Client(provider, region)
	if err != nil {
		return nil, err
	}
	os.Keystone, err = getIdentityClient(provider, region)
	if err != nil {
		return nil, err
	}

	log.Debug("OpenStack client initialized")
	return &os, nil
}

// getComputeClient 函数返回 ComputeV2 客户端
func getComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// getNetworkClient 函数返回 NetworkV2 客户端
func getNetworkClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// getImageClient 函数返回 ImageServiceV2 客户端
func getImageClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// getBlockStorageV2Client 函数返回 BlockStorageV2 客户端
func getBlockStorageV2Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// getBlockStorageV3Client 函数返回 BlockStorageV3 客户端
func getBlockStorageV3Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// getIdentityClient 函数返回 IdentityV3 客户端
func getIdentityClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}
