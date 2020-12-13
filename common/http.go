package common

import (
	"konggo/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusSuccess = 0 // 成功

	StatusServerError  = 2001 // 服务器内部错误
	StatusInvalidParam = 2002 // 参数错误
	StatusUnrealized   = 2003 // 未实现

	StatusTargetNotExist = 3001 //目标记录不存在
	StatusTargetIsExist  = 3002 //目标记录已存在
	StatusCantDelete     = 3003 //无法删除

	StatusNeedLogin = 4001 //未登录
)

const (
	//system
	MsgSuccess      = "成功"
	MsgServerError  = "服务器内部错误"
	MsgInvalidParam = "参数错误"
	MsgUnrealized   = "未实现"

	//database
	MsgTargetNotExist = "目标记录不存在"
	MsgTargetIsExist  = "同名记录已存在"
	MsgCantDelete     = "无法删除"

	//user
	MsgNeedLogin = "未登录"
)

//ResponseModel 响应数据,带参
type ResponseModel struct {
	Result    bool        `json:"result"`
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

// ResJson
func ResJson(c *gin.Context, data interface{}, err WebError) {

	if err == nil {
		ret := ResponseModel{Result: true, ErrorCode: StatusSuccess, Message: MsgSuccess, Data: data}
		ResJSON(c, http.StatusOK, &ret)
		return
	}

	ret := ResponseModel{Result: false, ErrorCode: err.Code(), Message: err.Msg(), Data: data}
	if err.Equal(nil) {
		ret.Result = true
	}

	ResJSON(c, http.StatusOK, &ret)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {

	logger.InfoKV("ResJSON", logger.KV{"status": status, "val": v})

	c.JSON(status, v)
	c.Abort()
}
