package v1

import (
	"autoCreate/models"
	"autoCreate/openstack"
	"autoCreate/utils"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func CreateVm(c *gin.Context) {
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
	var client *openstack.OpenStack
	// 创建 OpenStack 客户端
	client, err := openstack.NewOpenStack(conf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 获取镜像 ID
	imageID, err := client.GetImageId(request.ImageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get image ID"})
		return
	}

	// 获取云硬盘类型 ID
	volumeTypeID, err := client.GetVolumeTypeId(request.VolumeTypeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get volume type ID"})
		return
	}

	// 获取网络 ID
	networkIDs, err := client.GetNetworkIds(request.Networks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get network IDs"})
		return
	}

	// 获取云硬盘 ID
	volumeIDs, err := client.CreateVolumes(request.Volumes, volumeTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create volumes"})
		return
	}

	// 创建虚拟机
	createOpts := servers.CreateOpts{
		Name:      request.VmName,
		FlavorRef: request.FlavorName,
		ImageRef:  imageID,
		Networks:  networkIDs,
	}

	server, err := client.CreateInstance(createOpts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create virtual machine"})
		return
	}

	// 关联云硬盘
	for _, volumeID := range volumeIDs {
		err := client.AttachVolume(server, volumeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attach volume to virtual machine"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Virtual machine created successfully", "server": server})
}
