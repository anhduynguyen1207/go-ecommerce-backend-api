package repo

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/database"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user repo
// func (ur *UserRepo) GetInfoUser() string {
// 	return "Anh Duy User info"

// }

// INTERFACE
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func (up *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		fmt.Printf("GetUserByEmail error: %v\n", err)
		return false
	}
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
