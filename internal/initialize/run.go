package initialize

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Load config mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8888")
}
