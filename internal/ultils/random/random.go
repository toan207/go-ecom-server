package random

import (
	"math/rand"
	"strconv"
)

func GenerateSixLetterOTP() string {
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp)
}
