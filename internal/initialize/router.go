package initialize

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/routers"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/middlewares"
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

	r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 req/s
	r.GET("/ping/100", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 100",
		})
	})

	r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter()) // 80 req/s
	r.GET("/ping/80", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 80",
		})
	})

	r.Use(middlewares.NewRateLimiter().UserAndPrivateRateLimiter()) // 50 req/s
	r.GET("/ping/50", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 50",
		})
	})
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
