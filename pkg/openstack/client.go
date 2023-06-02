package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// OpenStack 结构体包含各个服务的客户端
type OpenStack struct {
	Nova     *gophercloud.ServiceClient
	Neutron  *gophercloud.ServiceClient
	Glance   *gophercloud.ServiceClient
	Cinder   CinderInterface
	Keystone *gophercloud.ServiceClient
}

// NewOpenStack 函数初始化 OpenStack 结构体
func NewOpenStack(conf OpenStackConfig) (*OpenStack, error) {
	provider, err := openstack.NewClient(conf.AuthURL)
	if err != nil {
		return nil, err
	}
	if err = openstack.Authenticate(provider, conf.AuthOpts()); err != nil {
		return nil, err
	}

	os := &OpenStack{}

	if os.Nova, err = os.NewComputeClient(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.Neutron, err = os.NewNetworkClient(provider, conf.Region); err != nil {
		return nil, err
	}

	if os.Glance, err = os.NewImageClient(provider, conf.Region); err != nil {
		return nil, err
	}

	cinder, version, err := os.NewBlockStorageClient(provider, conf.Region)
	if err != nil {
		return nil, err
	}
	os.Cinder = &Cinder{CinderServiceClient: cinder, Version: version}
	return os, nil
}

// NewComputeClient 函数返回 ComputeV2 客户端
func (os *OpenStack) NewComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewNetworkClient 函数返回 NetworkV2 客户端
func (os *OpenStack) NewNetworkClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewImageClient 函数返回 ImageServiceV2 客户端
func (os *OpenStack) NewImageClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV2Client 函数返回 BlockStorageV2 客户端
func (os *OpenStack) NewBlockStorageV2Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV3Client 函数返回 BlockStorageV3 客户端
func (os *OpenStack) NewBlockStorageV3Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

func (os *OpenStack) NewBlockStorageClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, int, error) {
	cinderv2, err := os.NewBlockStorageV3Client(provider, region)
	if err == nil {
		return cinderv2, 2, nil
	}

	cinderv3, err := os.NewBlockStorageV2Client(provider, region)
	if err == nil {
		return cinderv3, 3, nil
	}
	return nil, 0, fmt.Errorf("no available cinder service client")

}

// NewIdentityClient 函数返回 IdentityV3 客户端
func (os *OpenStack) NewIdentityClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}
