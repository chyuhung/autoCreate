package openstack

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// OpenStack 结构体包含各个服务的客户端
type OpenStack struct {
	Nova     *gophercloud.ServiceClient
	Neutron  *gophercloud.ServiceClient
	Glance   *gophercloud.ServiceClient
	Cinder   *CinderServiceClient
	Keystone *gophercloud.ServiceClient
}

// NewOpenStack 函数初始化 OpenStack 结构体
func NewOpenStack(conf OpenStackConfig) (*OpenStack, error) {
	log.Println(conf)
	provider, err := openstack.NewClient(conf.AuthURL)
	if err != nil {
		return nil, err
	}

	err = openstack.Authenticate(provider, conf.AuthOpts())
	if err != nil {
		return nil, err
	}

	nova, err := NewComputeClient(provider, conf.Region)
	if err != nil {
		return nil, err
	}

	neutron, err := NewNetworkClient(provider, conf.Region)
	if err != nil {
		return nil, err
	}

	glance, err := NewImageClient(provider, conf.Region)
	if err != nil {
		return nil, err
	}
	// 当前所有版本的Cinder
	var Cinder *CinderServiceClient
	cinderv1, errv1 := NewBlockStorageV1Client(provider, conf.Region)
	cinderv2, errv2 := NewBlockStorageV2Client(provider, conf.Region)
	cinderv3, errv3 := NewBlockStorageV3Client(provider, conf.Region)
	if errv1 == nil && errv2 == nil {
		Cinder = &CinderServiceClient{
			CinderServiceClientV1: &CinderServiceClientV1{
				CinderServiceClient: cinderv1,
				Version:             1},
			CinderServiceClientV2: &CinderServiceClientV2{
				CinderServiceClient: cinderv2,
				Version:             2},
		}
	}
	if errv2 == nil && errv3 == nil {
		Cinder = &CinderServiceClient{
			CinderServiceClientV3: &CinderServiceClientV3{
				CinderServiceClient: cinderv3,
				Version:             3},
			CinderServiceClientV2: &CinderServiceClientV2{
				CinderServiceClient: cinderv2,
				Version:             2},
		}
	}
	keystone, err := NewIdentityClient(provider, conf.Region)
	if err != nil {
		return nil, err
	}
	return &OpenStack{
		Nova:     nova,
		Neutron:  neutron,
		Cinder:   Cinder,
		Glance:   glance,
		Keystone: keystone,
	}, nil
}

// NewComputeClient 函数返回 ComputeV2 客户端
func NewComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewNetworkClient 函数返回 NetworkV2 客户端
func NewNetworkClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewImageClient 函数返回 ImageServiceV2 客户端
func NewImageClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV1Client 函数返回 BlockStorageV1 客户端
func NewBlockStorageV1Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV2Client 函数返回 BlockStorageV2 客户端
func NewBlockStorageV2Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewBlockStorageV3Client 函数返回 BlockStorageV3 客户端
func NewBlockStorageV3Client(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

// NewIdentityClient 函数返回 IdentityV3 客户端
func NewIdentityClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}
