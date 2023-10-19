package v1

import (
	"autoCreate/model"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddVm 添加虚拟机
func AddVm(c *gin.Context) {
	var data model.Vm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    data,
			"message": errmsg.JSON_ERROR,
		})
		return
	}
	code := model.CreateVm(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
