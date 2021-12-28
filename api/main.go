package api

import (
	"konggo/common"
	"konggo/model"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	common.ResJson(c, gin.H{"info": "Pong"}, nil)
}

// Version 编译版本确认页面
func Version(c *gin.Context) {
	common.ResJson(c, gin.H{"version": common.Version}, nil)
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}
