package errmsg

const (
	// 成功
	SUCCSE = 200
	// 失败
	ERROR = 500

	// others 5000...
	// nova 1000...
	// glance 2000...
	// cinder 3000...
	// neutron 4000...
	JSON_ERROR = "JSON 格式错误"

	// nova 1000...
	// glance 2000...
	// cinder 3000...
	// neutron 4000...

)

var codeMessage = map[int]string{
	SUCCSE: "OK",
	ERROR:  "FAIL",
}

// 获取错误信息
func GetErrMsg(code int) string {
	return codeMessage[code]
}
