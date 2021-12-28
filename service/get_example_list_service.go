package service

import (
	"konggo/common"
	"konggo/dao"
	"konggo/pkg/logger"
	"konggo/serializer"
)

// GetExampleListService 获取样例列表的服务
type GetExampleListService struct {
	common.Page
}

// GetExampleListService 样例列表查询的
func (service *GetExampleListService) GetExampleList() (res serializer.GetExampleList, e common.WebError) {

	total, examples, err := dao.GetExampleList(nil, service.Page)
	if err != nil {
		logger.Errorp(service, err)
		return serializer.GetExampleList{}, common.ErrServer()
	}

	return serializer.BuildExampleList(examples, total), nil
}
