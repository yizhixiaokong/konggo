package service

import (
	"konggo/common"
	"konggo/model"
	"konggo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"` //用户名
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`   //密码
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) (serializer.User, common.WebError) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.User{}, common.ErrInvalidParam().AddMsg(" 账号或密码错误")
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.User{}, common.ErrInvalidParam().AddMsg(" 账号或密码错误")
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUser(user), nil
}
