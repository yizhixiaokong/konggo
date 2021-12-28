package model

import "github.com/jinzhu/gorm"

// Example 样例模型
type Example struct {
	gorm.Model
	Name string
	Age  int
}

// TableName 表名
func (Example) TableName() string {
	return "example"
}
