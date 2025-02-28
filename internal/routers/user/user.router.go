package user

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	//this is non-depedency injection
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHanlderNonDependency := controller.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/opt")
	}

	//private
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_infor")
	}

}
