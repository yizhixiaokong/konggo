package api

import (
	"konggo/common"
	"konggo/pkg/logger"
	"konggo/serializer"
	"konggo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.Register()
	common.ResJson(c, res, err)

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.Login(c)
	common.ResJson(c, res, err)

}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUser(*user)
	common.ResJson(c, res, nil)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	common.ResJson(c, gin.H{"info": "登出成功"}, nil)
}
