package service

import (
	"konggo/common"
	"konggo/dao"
	"konggo/pkg/logger"
	"konggo/serializer"
)

// PutExampleService 修改样例
type PutExampleService struct {
	ID   uint   `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

// PutExample 样例修改的函数
func (service *PutExampleService) PutExample() (res serializer.Example, e common.WebError) {

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

	_, err = dao.GetExampleByName(nil, service.Name)
	exist, err = dao.ExistRow(err)
	if exist {
		return serializer.Example{}, common.ErrIsExist().AddMsg(" :样例name已存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Example{}, common.ErrServer()
	}

	example.Name = service.Name
	example.Age = service.Age

	err = dao.TxSave(nil, &example)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.Example{}, common.ErrServer()
	}

	return serializer.BuildExample(example), nil
}
