package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/repo"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/utils/random"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/anhduynguyen1207/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// Interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}
func (us *userService) Register(email string, purpose string) int {
	//0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail is :::%s\n", hashEmail)
	// 5. Check OPT is available
	// 6. Check use spam ...
	// 1. Check email exist in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	// 2. new OTP -> ....
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is :::%d\n", otp)
	// 3. save OTP in Redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	// fmt.Printf("err is :::%d\n", err)
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. send email OPT
	// err = sendto.SendTemplateEmailOtp([]string{email}, "anonystick@gmail.com", "otp-auth.html"), map[string]interface{}
	// err = sendto.SendTemplateEmailOtp([]string{email}, "anonystick@gmail.com", "otp-auth.html", map[string]interface{}{
	// 	"otp": strconv.Itoa(otp),
	// })
	// // fmt.Printf("err sendto :::%d\n", err)
	// if err != nil {
	// 	return response.ErrSendEmailOtp
	// }
	// return response.ErrCodeSuccess

	// send email OTP by JAVA
	err = sendto.SendEmailToJavaByApi(strconv.Itoa(otp), email, "otp-auth.html")
	// fmt.Printf("err sendto :JAVA::%d\n", err)
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}
