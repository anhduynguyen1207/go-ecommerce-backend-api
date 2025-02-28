package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	//public router
	ProductRouterPublic := Router.Group("/product")
	{
		ProductRouterPublic.GET("/search")
		ProductRouterPublic.GET("/detail/:id")
	}

	//private

}
