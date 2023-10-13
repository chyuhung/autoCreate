package routes

import (
	v1 "autoCreate/api/v1"
	"autoCreate/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	auth := r.Group("api/v1")
	{
		// 登录模块
		auth.POST("/login", v1.LoginHandler)

		// Nova 模块路由接口
		// 所有实例列表
		auth.GET("/instances", v1.GetInstances)

		// Glance 模块路由接口
		// 所有镜像列表
		auth.GET("/images", v1.GetImages)

		// Cinder 模块路由接口
		// 所有卷列表
		auth.GET("/volumes", v1.GetVolumes)

		// Neutron 模块路由接口
		// 所有网络列表
		auth.GET("/networks", v1.GetNetworks)
	}
	r.Run(utils.HttpPort)
}
