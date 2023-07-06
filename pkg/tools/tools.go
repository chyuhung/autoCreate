package tools

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sahilm/fuzzy"
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
	if len(array) == 0 {
		return -1, fmt.Errorf("no valid elements found")
	}
	if len(array) == 1 {
		return 0, nil
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
	lines := SplitWithoutEmpty(string(data), '\n')
	return lines, nil
}
func SplitWithoutEmpty(s string, sep rune) []string {
	parts := strings.FieldsFunc(s, func(c rune) bool {
		return c == sep
	})
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if !strings.ContainsAny(part, string(sep)) {
			result = append(result, part)
		}
	}
	return result
}

// 写入数据到文件
func WriteToEnvFile(array []string, filename string) error {
	if len(array) > 0 {
		data := strings.Join(array, "\n") + "\n"
		return os.WriteFile(filename, []byte(data), 0644)
	}
	return fmt.Errorf("empty array, do not write")
}

// 去重 string 数组
func UniqueStrings(s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	var result []string
	for key := range m {
		result = append(result, key)
	}
	return result
}

// 模糊匹配， 传入字符串和待匹配切片/数组，返回匹配到的最佳的切片/数组中的字符串和错误
func FuzzyMatch(input string, options []string) (string, error) {
	matches := fuzzy.Find(input, options)
	if len(matches) == 0 {
		return input, fmt.Errorf("no matching string found")
	}
	return matches[0].Str, nil
}
