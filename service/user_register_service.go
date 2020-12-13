package service

import (
	"konggo/common"
	"konggo/model"
	"konggo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`                 //昵称
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`               //用户名
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`                 //密码
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"` //确认密码
}

// valid 验证表单
func (service *UserRegisterService) valid() (err common.WebError) {
	if service.PasswordConfirm != service.Password {
		return common.ErrInvalidParam().AddMsg(" 两次密码不一致")
	}
	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return common.ErrIsExist().AddMsg(" 昵称被占用")
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return common.ErrIsExist().AddMsg(" 用户名已经注册")
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (serializer.User, common.WebError) {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return serializer.User{}, err
	}

	// 加密密码
	if e := user.SetPassword(service.Password); e != nil {
		return serializer.User{}, common.ErrServer().AddMsg(" 密码加密失败 ").AddMsg(e.Error())
	}

	// 创建用户
	if e := model.DB.Create(&user).Error; e != nil {
		return serializer.User{}, common.ErrServer().AddMsg(" 注册失败 ").AddMsg(e.Error())
	}

	return serializer.BuildUser(user), nil

}
