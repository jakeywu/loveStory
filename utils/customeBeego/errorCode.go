package customeBeego

// 自定义错误日志
const (
	// 400
	ErrorParameter = 400101
	ErrorValidate  = 400102

	// 403
	ErrorAuthentication = 403001

	// 404
	ErrorNotFound = 404001

	// 500
	ErrorValidServer = 503001
)

var errorText = map[int]string{
	ErrorParameter: "数据解析失败",
	ErrorValidate:  "用户名或密码格式有误",

	ErrorAuthentication: "用户名或密码错误",

	ErrorNotFound: "资源不存在",

	// valid　struct有误, 请检查
	ErrorValidServer: "系统错误, 请联系管理员",
}

func ErrorText(code int) string {
	return errorText[code]
}
