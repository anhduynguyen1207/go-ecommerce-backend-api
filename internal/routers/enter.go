package routers

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
