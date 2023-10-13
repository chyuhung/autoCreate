package v1

import (
	"github.com/gin-gonic/gin"
)

func GetInstances(c *gin.Context) {
	// code := errmsg.SUCCSE
	// pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	// marker := c.Query("marker")

	// data, err := openstack.GetInstances(pageSize, marker)
	// if err != nil {
	// 	code = errmsg.ERROR
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  code,
	// 	"data":    data,
	// 	"total":   len(data),
	// 	"message": err,
	// })
}
