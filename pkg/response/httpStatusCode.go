package response

// những code này sẽ tự qui định trong nội bộ công ty BE và FE. Và từ đó FE dựa vào ErrorCode để tiến hành xử lý
// Ví dụ như: ErrorCode = 20003 -> Email không hợp lệ
const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 //Email is invalid

	ErrInvalidToken = 30001 // Token không hợp lệ
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
}
