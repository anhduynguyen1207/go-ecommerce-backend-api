package user

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/controller/account"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/wire"
	middleware "github.com/anhduynguyen1207/go-ecommerce-backend-api/middlewares"
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
		// userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/verify_update_pass_register", account.Login.UpdatePasswordRegister)

		// userRouterPublic.POST("/opt")
		userRouterPublic.POST("/login", account.Login.Login)

	}

	//private
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middleware.AuthenMiddleware())
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info", userController.Register)
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	}

}
