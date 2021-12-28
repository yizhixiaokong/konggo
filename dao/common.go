package dao

import (
	"database/sql"
	database "konggo/model"

	"github.com/jinzhu/gorm"
)

//PageWhereOrder 分页条件
type PageWhereOrder struct {
	//order
	Order string
	//where
	Where string
	Value []interface{}
}

// Create ...
func Create(value interface{}) error {
	return database.GetDB().Create(value).Error
}

// Create ...
func TxCreate(tx *gorm.DB, value interface{}) error {
	if tx == nil {
		tx = database.GetDB()
	}
	return tx.Create(value).Error
}

// Save 更新所有字段无论是否为空
func Save(value interface{}) error {
	return database.GetDB().Save(value).Error
}

func TxSave(tx *gorm.DB, value interface{}) error {
	if tx == nil {
		tx = database.GetDB()
	}
	return tx.Save(value).Error
}

// Updates ...
// 不会更新空字段如 "" , 0 , false
func Updates(model interface{}, new interface{}) error {
	return database.GetDB().Model(model).Updates(new).Error
}

// DeleteByModel Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	db := database.GetDB().Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByWhere Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := database.GetDB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID Delete
func DeleteByID(model interface{}, id int64) (count int64, err error) {
	db := database.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByIDS Delete
func DeleteByIDS(model interface{}, ids []int64) (count int64, err error) {
	db := database.GetDB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// FirstByID First
func FirstByID(out interface{}, id int64) (notFound bool, err error) {
	err = database.GetDB().First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First ...
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = database.GetDB().Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find ...
func Find(where interface{}, out interface{}, orders ...string) error {
	db := database.GetDB().Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan ...
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = database.GetDB().Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList ...
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := database.GetDB().Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// GetPage ...
func GetPage(model interface{}, where interface{}, out interface{}, pageIndex, pageSize int, totalCount *int, whereOrder ...PageWhereOrder) error {
	db := database.GetDB().Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList ...
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return database.GetDB().Model(model).Where(where).Pluck(fieldName, out).Error
}

// ExistRow 查询存在
func ExistRow(err error) (bool, error) {
	if err == sql.ErrNoRows {
		return false, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}

	return true, err
}

// TxBegin 事务开启
func TxBegin() (*gorm.DB, error) {
	tx := database.GetDB().Begin()
	return tx, tx.Error
}
