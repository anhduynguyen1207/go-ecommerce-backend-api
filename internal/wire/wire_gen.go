// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/controller"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/repo"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserAuthRepository := repo.NewUserAuthRepository()
	iUserService := service.NewUserService(iUserRepository, iUserAuthRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
