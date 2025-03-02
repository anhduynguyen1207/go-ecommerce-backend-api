package response

// những code này sẽ tự qui định trong nội bộ công ty BE và FE. Và từ đó FE dựa vào ErrorCode để tiến hành xử lý
// Ví dụ như: ErrorCode = 20003 -> Email không hợp lệ
const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 //Email is invalid

	ErrInvalidToken = 30001 // Token không hợp lệ
	ErrInvalidOTP   = 30002 // OTP không hợp lệ
	ErrSendEmailOtp = 30003 // Gửi email OTP thất bại
	//Register Code
	ErrCodeUserHasExists = 50001 // User has already exists

	//Err Login
	ErrCodeOtpNotExists    = 60009 // OTP không tồn tại nhưng không đăng ký
	ErrCodeUserOtpNotExist = 60008 //OTP không tồn tại
	//User AUthentication
	ErrCodeLoginFail = 40005 // Login fail
)

// message
var msg = map[int]string{
	ErrCodeSuccess:         "Success",
	ErrCodeParamInvalid:    "Email is invalid",
	ErrInvalidToken:        "Token is invalid",
	ErrInvalidOTP:          "OTP is invalid",
	ErrSendEmailOtp:        "Fail to send email OTP",
	ErrCodeUserHasExists:   "User has already exists",
	ErrCodeOtpNotExists:    "OTP is not exists but not registed",
	ErrCodeUserOtpNotExist: "OTP is not exists",
	ErrCodeLoginFail:       "Login fail",
}
