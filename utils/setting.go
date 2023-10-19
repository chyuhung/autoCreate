package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	HttpPort    string
	Username    string
	Password    string
	ProjectName string
	DomainName  string
	Region      string
	AuthURL     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误:", err)
	}
	loadServer(file)
}

func loadServer(file *ini.File) {
	// server
	sectionServer := file.Section("server")
	HttpPort = sectionServer.Key("HttpPort").MustString("3000")

	// openstack
	sectionOpenstack := file.Section("openstack")
	Username = sectionOpenstack.Key("OS_USERNAME").MustString("admin")
	Password = sectionOpenstack.Key("OS_PASSWORD").MustString("")
	ProjectName = sectionOpenstack.Key("OS_PROJECT_NAME").MustString("admin")
	DomainName = sectionOpenstack.Key("OS_USER_DOMAIN_NAME").MustString("Default")
	Region = sectionOpenstack.Key("OS_REGION_NAME").MustString("RegionOne")
	AuthURL = sectionOpenstack.Key("OS_AUTH_URL").MustString("")

	if Username == "" || Password == "" || ProjectName == "" || DomainName == "" || AuthURL == "" {
		log.Panicln("failed to read valid OS_USERNAME, OS_PASSWORD, OS_PROJECT_NAME, OS_USER_DOMAIN_NAME, or OS_AUTH_URL from the configuration file")
		os.Exit(1)
	}
}
