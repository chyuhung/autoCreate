package v1

import (
	"autoCreate/models"
	"autoCreate/openstack"
	"autoCreate/utils"
	"autoCreate/utils/errmsg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建虚拟机
func CreateVm(c *gin.Context) {
	// 获取请求参数
	var request models.VmRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errmsg.JSON_ERROR})
		return
	}

	// 配置 OpenStack 认证信息
	conf := openstack.OpenStackConfig{
		Username:    utils.Username,
		Password:    utils.Password,
		ProjectName: utils.ProjectName,
		DomainName:  utils.DomainName,
		AuthURL:     utils.AuthURL,
		Region:      utils.Region,
	}
	log.Println("Openstack配置信息:", conf)
	// var client *openstack.OpenStack
	// // 创建 OpenStack 客户端
	// client, err := openstack.NewOpenStack(conf)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	// // 获取镜像 ID
	// imageID, err := client.GetImageId(request.ImageName)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	// // 获取云硬盘类型 ID
	// volumeTypeID, err := client.GetVolumeTypeId(request.VolumeTypeName)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get volume type Id"})
	// 	return
	// }

	// // 获取网络
	// networks, err := client.BuildNetworks(request.Networks)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get network Ids"})
	// 	return
	// }

	// // 获取云硬盘 ID
	// volumeIDs, err := client.CreateVolumes(request.Volumes, volumeTypeID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create volumes"})
	// 	return
	// }

	// // 创建虚拟机
	// createOpts := servers.CreateOpts{
	// 	// server name
	// 	Name: request.VmName,
	// 	// flavor id
	// 	FlavorRef: request.FlavorName,
	// 	// image id
	// 	ImageRef: imageID,
	// 	Networks: networks,
	// }

	// server, err := client.CreateInstance(createOpts)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// // 关联云硬盘
	// for _, volumeID := range volumeIDs {
	// 	err := client.AttachVolume(server, volumeID)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attach volume to virtual machine"})
	// 		return
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Virtual machine created successfully", "server": "server"})
}
