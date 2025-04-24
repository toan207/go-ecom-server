package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20003
)

var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
}
