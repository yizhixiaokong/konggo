package serializer

import "konggo/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`         //id
	UserName  string `json:"user_name"`  //用户名
	Nickname  string `json:"nickname"`   //昵称
	Status    string `json:"status"`     //状态
	Avatar    string `json:"avatar"`     //头像
	CreatedAt int64  `json:"created_at"` //创建时间
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
