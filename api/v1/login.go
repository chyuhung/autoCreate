package v1

import (
	"autoCreate/openstack"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var conf openstack.OpenStackConfig
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"data":   errmsg.JSON_ERROR,
		})
		return
	}
	token := generateToken()
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCSE,
		"token":  token,
		"data":   conf,
	})
}

// 获取token
func generateToken() string {
	return "Login successful."
}
