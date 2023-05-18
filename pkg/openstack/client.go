package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// OpenStack 结构体包含各个服务的客户端
type openStack struct {
	Nova     *gophercloud.ServiceClient
	Neutron  *gophercloud.ServiceClient
	Glance   *gophercloud.ServiceClient
	CinderV2 *gophercloud.ServiceClient
	CinderV3 *gophercloud.ServiceClient
	Keystone *gophercloud.ServiceClient
}

// NewOpenStack 函数初始化 OpenStack 结构体
func NewOpenStack(conf OpenStackConfig) (*openStack, error) {
	provider, err := openstack.NewClient(conf.AuthURL)
	if err != nil {
		return nil, err
	}
	if err = openstack.Authenticate(provider, conf.AuthOpts()); err != nil {
		return nil, err
	}

	os := &openStack{}

	if os.Nova, err = os.NewComputeClient(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.Neutron, err = os.NewNetworkClient(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.Glance, err = os.NewImageClient(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.CinderV2, err = os.NewBlockStorageV2Client(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.CinderV3, err = os.NewBlockStorageV3Client(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.Keystone, err = os.NewIdentityClient(provider, conf.Region); err != nil {
		return nil, err
	}

	//log.Info("openstack client initialized")
	return os, nil
}

// NewComputeClient 函数返回 ComputeV2 客户端
func (os *openStack) NewComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewNetworkClient 函数返回 NetworkV2 客户端
func (os *openStack) NewNetworkClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewImageClient 函数返回 ImageServiceV2 客户端
func (os *openStack) NewImageClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV2Client 函数返回 BlockStorageV2 客户端
func (os *openStack) NewBlockStorageV2Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV3Client 函数返回 BlockStorageV3 客户端
func (os *openStack) NewBlockStorageV3Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewIdentityClient 函数返回 IdentityV3 客户端
func (os *openStack) NewIdentityClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}
