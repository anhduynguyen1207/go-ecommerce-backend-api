package controller

import (
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "No need")
	// }
	// return response.SuccessResponse(c, 20001, []string{"anhduy", "quocanh", "mimi"})

}

// luồng dữ liệu: controller -> service -> repo -> models -> db
