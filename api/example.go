package api

import (
	"konggo/common"
	"konggo/pkg/logger"
	"konggo/service"

	"github.com/gin-gonic/gin"
)

// 获取样例列表
func GetExampleList(c *gin.Context) {
	var service service.GetExampleListService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.GetExampleList()
	common.ResJson(c, res, err)
}

// 获取样例
func GetExample(c *gin.Context) {
	var service service.GetExampleService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.GetExample()
	common.ResJson(c, res, err)
}

// 添加样例
func PostExample(c *gin.Context) {
	var service service.PostExampleService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.PostExample()
	common.ResJson(c, res, err)
}

// 修改样例
func PutExample(c *gin.Context) {
	var service service.PutExampleService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	res, err := service.PutExample()
	common.ResJson(c, res, err)
}

// 删除样例
func DeleteExample(c *gin.Context) {
	var service service.DeleteExampleService
	if err := c.ShouldBind(&service); err != nil {
		logger.Info(err.Error())
		common.ResJson(c, nil, common.ErrInvalidParams(err))
		return
	}
	err := service.DeleteExample()
	common.ResJson(c, nil, err)
}
