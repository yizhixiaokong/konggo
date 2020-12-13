package conf

import (
	"fmt"
	"konggo/cache"
	"konggo/model"
	"konggo/pkg/logger"
	"konggo/util"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//main
var (
	ApplicationPath string //应用程序路径
	HTTPServer      string //文件存储服务器地址
)

// Init 初始化配置项
func Init() {
	// 运行目录
	initAppPath()

	// 从本地读取环境变量
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		panic("godotenv failed")
	}

	// 初始化日志系统
	s := os.Getenv("LOG_SECRET")
	encrypted := true
	if s == "encoding" {
		encrypted = false
	}
	logger.InitLog(ApplicationPath, encrypted)

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}

func initAppPath() {
	//初始化 ApplicationPath
	if file, err := os.Getwd(); err != nil {
		panic("init ApplicationPath failed")
	} else {
		ApplicationPath = file + string(os.PathSeparator)
	}
	log.Printf("\n========ApplicationPath========\n%s\n", ApplicationPath)
}
