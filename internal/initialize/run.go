package initialize

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Load config mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("ok", "Logger init success"))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8888")
}
