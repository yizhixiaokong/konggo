package dao

import (
	"konggo/common"
	"konggo/model"

	"github.com/jinzhu/gorm"
)

func GetExampleList(tx *gorm.DB, page common.Page) (total int, items []model.Example, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	name := model.Example{}.TableName()
	tx = tx.Table(name)

	err = tx.Where("`deleted_at` IS NULL").Count(&total).Error //total
	if total == 0 {
		return total, items, err
	}

	//分页
	tx = tx.
		Order(page.Order("id")).
		Offset(page.Offset()).
		Limit(page.Limit())

	err = tx.Find(&items).Error

	return
}

//通过id查找
func GetExample(tx *gorm.DB, id uint) (item model.Example, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	tableName := model.Example{}.TableName()
	tx = tx.Table(tableName)

	where := map[string]interface{}{"id": id}

	err = tx.Where(where).First(&item).Error
	return
}

//通过name查找
func GetExampleByName(tx *gorm.DB, name string) (item model.Example, err error) {
	if tx == nil {
		tx = model.GetDB()
	}

	tableName := model.Example{}.TableName()
	tx = tx.Table(tableName)

	where := map[string]interface{}{"name": name}

	err = tx.Where(where).First(&item).Error
	return
}
