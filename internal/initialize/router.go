package initialize

import (
	"fmt"

	c "github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/controller"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/middleware"
	"github.com/gin-gonic/gin"
)

// middleware
func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")

}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthenMiddleware(), AA(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong) //v1/2024/ping
		v1.GET("/user/1", c.NewUserController().GetUserByID)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	// v2 := r.Group("/v2/2024")
	// {
	// v2.GET("/ping", Pong) //v2/2024/ping
	// v2.PUT("/ping", Pong)
	// v2.PATCH("/ping", Pong)
	// v2.DELETE("/ping", Pong)
	// v2.HEAD("/ping", Pong)
	// v2.OPTIONS("/ping", Pong)
	// }

	return r
}
