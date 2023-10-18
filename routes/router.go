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
	}
	r.Run(utils.HttpPort)
}
