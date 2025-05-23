package response

const (
	ErrorCodeSuccess        = 20001
	ErrorCodeParamInvalid   = 20003
	ErrorCodeOTPError       = 30001 // OTP is invalid
	ErrorCodeInvalidParams  = 30002 // Invalid parameters
	ErrorCodeHashEmailError = 40001 // Hash email error
	ErrorCodeSendOTPErr     = 40002 // Send OTP error
	ErrorCodeEmailExist     = 50001 // Email already exists
	ErrorCodeInternalServer = 50002 // Internal server error
)

var msg = map[int]string{
	ErrorCodeSuccess:        "Success",
	ErrorCodeParamInvalid:   "Email is invalid",
	ErrorCodeOTPError:       "OTP error",
	ErrorCodeHashEmailError: "Hash email error",
	ErrorCodeEmailExist:     "Email already exists",
	ErrorCodeInternalServer: "Internal server error",
	ErrorCodeSendOTPErr:     "Send OTP error",
	ErrorCodeInvalidParams:  "Invalid parameters",
}
