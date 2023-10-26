// Initialize router
package routes

import (
	v1 "autoCreate/controllers/v1"
	"autoCreate/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	auth := r.Group("api/v1")
	{
		// 登录模块
		auth.POST("/login", v1.LoginHandler)

		// 创建虚拟机
		auth.POST("/createVm", v1.CreateVm)

	}
	r.Run(utils.HttpPort)
}
