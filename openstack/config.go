package openstack

import "github.com/gophercloud/gophercloud"

type OpenStackConfig struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ProjectName string `json:"project_name"`
	DomainName  string `json:"domain_name"`

	AuthURL string `json:"auth_url"`
	Region  string `json:"region"  default:"ReginOne"`
}

// AuthOpts gets openstack auth options
func (conf OpenStackConfig) AuthOpts() gophercloud.AuthOptions {
	return gophercloud.AuthOptions{
		IdentityEndpoint: conf.AuthURL,
		Username:         conf.Username,
		Password:         conf.Password,
		DomainName:       conf.DomainName,
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: conf.ProjectName,
			DomainName:  conf.DomainName},
	}
}
