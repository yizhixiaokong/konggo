package model

import (
	"konggo/pkg/logger"
)

// GormLogger struct
type GormLogger struct{}

//Print 实现logger接口
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logger.Infof("[gorm sql]:\n<src>: (%v)\n<duration>: [%v]\n<sql>: [%v]\n<values>: [%v]\n<rows_returned>: [%v rows affected or returned] \n", v[1], v[2], v[3], v[4], v[5])
	case "log":
		logger.Infof("[gorm log]: %v\n", v[2])
	}
}
