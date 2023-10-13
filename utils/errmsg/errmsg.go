package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// others 5000...
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

func GetErrMsg(code int) string {
	return codeMessage[code]
}
