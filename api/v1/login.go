package v1

import (
	"autoCreate/openstack"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// user conf
	var conf openstack.OpenStackConfig
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ERROR,
			"message": err,
		})
	}
	token := generateToken()
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCSE,
		"data":   token,
	})
}

// 获取token
func generateToken() string {
	return "test success"
}
