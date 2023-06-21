package cmd

import (
	"encoding/csv"
	"os"
	"strings"
)

// 存储从 csv 文件中读入的 instance info 信息
type vmInfo struct {
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

func readCsvFile(file string) ([]vmInfo, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var vmInfoArray []vmInfo
	for i, record := range records {
		if i == 0 {
			continue // skip header row
		}
		var instance vmInfo
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
		vmInfoArray = append(vmInfoArray, instance)
	}

	// Use the instances
	// for _, instance := range instances {
	// 	fmt.Printf("%+v\n", instance)
	// }
	return vmInfoArray, nil
}
