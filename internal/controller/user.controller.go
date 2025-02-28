package controller

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/service"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/vo"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Printf("Email params: %s", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// if err != nil {
// 	// 	return response.ErrorResponse(c, 20003, "No need")
// 	// }
// 	// return response.SuccessResponse(c, 20001, []string{"anhduy", "quocanh", "mimi"})

// }

// luồng dữ liệu: controller -> service -> repo -> models -> db
