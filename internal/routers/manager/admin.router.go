package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	//private
	adminRouterPrivate := Router.Group("/admin/user")
	// AdminRouterPrivate.Use(limiter())
	// AdminRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_admin")
	}

}
