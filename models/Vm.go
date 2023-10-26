package models

import (
	"autoCreate/utils/errmsg"
	"fmt"

	"gorm.io/gorm"
)

// 定义一个VmRequest结构体
type VmRequest struct {
	gorm.Model
	ImageName      string            `label:"镜像名称" json:"image_name"`
	FlavorName     string            `label:"规格名称" json:"flavor_name"`
	Volumes        map[string]int    `label:"卷" json:"volumes"`
	VmName         string            `label:"虚拟机名称" json:"vm_name"`
	Networks       map[string]string `label:"网络" json:"networks"`
	HostName       string            `label:"宿主机名称" json:"host_name"`
	VolumeTypeName string            `label:"卷类型名称" json:"volume_type_name"`
}

// 创建虚拟机
func CreateVm(data *VmRequest) int {
	// 创建卷
	// 创建虚拟机
	return errmsg.SUCCSE
}

// 创建卷
func createVol(volumeName string, volumeSize int, volumeType string) error {
	return fmt.Errorf("test")
}
