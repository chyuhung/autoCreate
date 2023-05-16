package tools

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 获取IP的最后一位数
func GetIpLastNum(ip string) int {
	lastDotIndex := strings.LastIndex(ip, ".")
	lastByte, _ := strconv.Atoi(ip[lastDotIndex+1:])
	return lastByte
}
func GetIpLastNu(ip string) int {
	parts := strings.Split(ip, ".")
	// parts 最后一段 ，parts最后一段的长度-1
	lastByte := string(parts[len(parts)-1][len(parts[len(parts)-1])-1])
	lastNum, _ := strconv.Atoi(lastByte)
	// fmt.Printf("last num is %d\n", lastNum)
	return lastNum
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// 随机获取array下标
func GetRandIndex(array []string) (int, error) {
	if len(array) < 1 {
		return -1, fmt.Errorf("array length error")
	}
	return r.Intn(len(array)), nil
}

// 删除特殊字符,空格除外
func RemoveSpecialChars(line string) string {
	reg, _ := regexp.Compile(`[^a-zA-Z0-9\s]+`)
	return reg.ReplaceAllString(line, "")
}

// 添加“,”符号
func AddSpecialChar(line string) string {
	return line + ","
}

// 读取文件内容
func ReadFromEnvFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

// 写入数据到文件
func WriteToEnvFile(array []string, filename string) error {
	return os.WriteFile(filename, []byte(strings.Join(array, "\n")), 0644)
}