package service

import (
	"konggo/common"
	"konggo/dao"
	"konggo/pkg/logger"
)

// DeleteExampleService 删除样例
type DeleteExampleService struct {
	ID uint `form:"id" json:"id"`
}

// DeleteExample 样例删除的函数
func (service *DeleteExampleService) DeleteExample() (e common.WebError) {

	var exist bool
	example, err := dao.GetExample(nil, service.ID)
	exist, err = dao.ExistRow(err)
	if !exist {
		return common.ErrNotExist().AddMsg(" :样例id不存在")
	}
	if err != nil {
		logger.Errorp(service, err)
		return common.ErrServer()
	}

	_, err = dao.DeleteByID(&example, int64(example.ID))
	if err != nil {
		logger.Errorp(service, err)
		return common.ErrServer()
	}

	return nil
}
