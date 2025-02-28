//go:build wireinject

package wire

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/controller"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/repo"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
