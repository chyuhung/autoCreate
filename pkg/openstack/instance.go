package openstack

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// InstanceInfo csv文件包含的实例信息
type InstanceInfo struct {
	Name        string
	OsName      string
	Cpu         string
	Mem         string
	SysVolSize  string
	DataVolSize string
	VlanName    string
	Ipaddr      string
	// 多ip支持
	Networks map[string]string // eg: vlan 2032:10.191.22.45
}

// 读取csv文件
func ReadCsvFile() []InstanceInfo {
	// Compile the regular expression pattern
	pattern := regexp.MustCompile(`^openstack.*\.csv$`)

	// Read all files in the current directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err.Error())
	}

	// Find the matching files
	var matchingFiles []os.FileInfo
	for _, file := range files {
		if pattern.MatchString(file.Name()) {
			matchingFiles = append(matchingFiles, file)
		}
	}

	// Create an array of instances
	var instances []InstanceInfo

	// Read the matching files
	for _, file := range matchingFiles {
		// Open the file
		f, err := os.Open(file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer f.Close()

		// Read the CSV records
		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i, record := range records {
			if i == 0 {
				continue // skip header row
			}
			var instance InstanceInfo
			// 11
			instance.Name = record[11]
			// 6
			instance.OsName = record[6]
			// 7
			instance.Cpu = record[7]
			// 8
			instance.Mem = record[8]
			// 9
			instance.SysVolSize = record[9]
			// 10
			instance.DataVolSize = record[10]
			// 12
			instance.VlanName = record[12]
			// 13
			instance.Ipaddr = record[13]
			// 14
			instance.Networks = make(map[string]string)
			// 如果存在多ip
			if len(record[14]) > 0 {
				networks := strings.Split(record[14], ";")
				for j := 0; j < len(networks); j++ {
					// Parse the network information
					parts := strings.Split(networks[j], ":")
					instance.Networks[parts[0]] = parts[1]
					//fmt.Println("networks[", j, "]:", networks[j])
				}
			}
			instances = append(instances, instance)
		}

		// Use the instances
		// for _, instance := range instances {
		// 	fmt.Printf("%+v\n", instance)
		// }
	}
	return instances
}
