package initialize

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("ok", "Logger init success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	r := InitRouter()
	return r
	// r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
