package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20003
	ErrorCodeEmailExist   = 50001 // Email already exists
)

var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorCodeEmailExist:   "Email already exists",
}
