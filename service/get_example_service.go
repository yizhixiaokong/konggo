package service

import (
	"konggo/common"
	"konggo/dao"
	"konggo/pkg/logger"
	"konggo/serializer"
)

// GetExampleService 获取样例
type GetExampleService struct {
	ID uint `form:"id" json:"id"`
}

// GetExample 样例查询的函数
func (service *GetExampleService) GetExample() (res serializer.Example, e common.WebError) {

	var exist bool
	example, err := dao.GetExample(nil, service.ID)
	exist, err = dao.ExistRow(err)
	if !exist {
		return serializer.Example{}, common.ErrNotExist().AddMsg(" :样例id不存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Example{}, common.ErrServer()
	}

	return serializer.BuildExample(example), nil
}
