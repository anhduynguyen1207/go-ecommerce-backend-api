package initialize

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/routers"
	"github.com/gin-gonic/gin"
)

// middleware

func InitRouter() *gin.Engine {
	// r := gin.Default()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	// r.Use() //logging
	// r.Use() //cross
	// r.Use() //limiter global
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)

	}
	return r
}
