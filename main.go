package main

import (
	"autoCreate/openstack"
	"autoCreate/routes"
	"autoCreate/utils"
	"log"
)

var CONF openstack.OpenStackConfig

func main() {
	// 读取配置
	CONF = openstack.OpenStackConfig{
		Username:    utils.Username,
		Password:    utils.Password,
		ProjectName: utils.ProjectName,
		DomainName:  utils.DomainName,
		AuthURL:     utils.AuthURL,
		Region:      utils.Region,
	}
	// 获取openstack客户端
	client, err := openstack.NewOpenStack(CONF)
	if err != nil {
		log.Panicln("连接openstack失败")
	}
	routes.InitRouter()
}
