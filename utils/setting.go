package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	HttpPort string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误:", err)
	}
	loadServer(file)
}

func loadServer(file *ini.File) {
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}
