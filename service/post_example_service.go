package service

import (
	"konggo/common"
	"konggo/dao"
	"konggo/model"
	"konggo/pkg/logger"
	"konggo/serializer"
)

// PostExampleService 添加样例
type PostExampleService struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

// PostExample 样例添加的函数
func (service *PostExampleService) PostExample() (res serializer.Example, e common.WebError) {

	var exist bool
	_, err := dao.GetExampleByName(nil, service.Name)
	exist, err = dao.ExistRow(err)
	if exist {
		return serializer.Example{}, common.ErrIsExist().AddMsg(" :样例name已存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Example{}, common.ErrServer()
	}

	example := model.Example{
		Name: service.Name,
		Age:  service.Age,
	}

	err = dao.TxCreate(nil, &example)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Example{}, common.ErrServer()
	}

	return serializer.BuildExample(example), nil
}
