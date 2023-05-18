package openstack

import "github.com/gophercloud/gophercloud"

type OpenStackConfig struct {
	Username    string
	Password    string
	ProjectName string
	DomainName  string

	AuthURL string
	Region  string
}

// AuthOpts gets openstack auth options
func (conf OpenStackConfig) AuthOpts() gophercloud.AuthOptions {
	return gophercloud.AuthOptions{
		IdentityEndpoint: conf.AuthURL,
		Username:         conf.Username,
		Password:         conf.Password,
		DomainName:       conf.DomainName,
		AllowReauth:      true,
	}
}
