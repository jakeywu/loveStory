package customeBeego

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"sc-git.com/beeApi/utils"
	"strconv"
	"strings"
)

type CustomController struct {
	beego.Controller
}

type CustomResponse struct {
	ReqId   string      `json:"requestId"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 统一初始化
func (c *CustomController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ct, controllerName, actionName, app)
	if strings.ToUpper(c.Ctx.Request.Method) != "GET" && strings.Contains(c.Ctx.Request.Header.Get("Content-Type"), "application/json") {
		c.ErrorHandler(ErrorParameter, "请求头类型有误")
		return
	}
}

// 自定义错误类型
func (c *CustomController) ErrorHandler(errCode int, errMsg... string) {
	httpCode, err := strconv.Atoi(strconv.Itoa(errCode)[0:3])
	if err != nil {
		panic("自定义代码错误")
	}
	c.Ctx.Output.SetStatus(httpCode)
	cr := new(CustomResponse)
	cr.ReqId = utils.UUID()
	cr.Code = errCode
	if len(errMsg) == 0 {
		cr.Message = ErrorText(errCode)
	} else {
		cr.Message = errMsg[0]
	}
	c.Data["json"] = cr
	c.ServeJSON()
}

// 自定义返回的Json结构
func (c *CustomController) CustomJsonResponse(result interface{}) {
	cr := new(CustomResponse)
	cr.Data = result
	cr.Code = 200
	cr.Message = "SUCCESS"
	cr.ReqId = utils.UUID()
	c.Data["json"] = cr
	c.ServeJSON()
}
