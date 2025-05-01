package repo

import (
	"fmt"
	"time"

	"pawtopia.com/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp string, expiredTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp string, expiredTime int64) error {
	key := fmt.Sprintf("usr::%s::otp::", email)
	return global.Redis.SetEx(ctx, key, otp, time.Duration(expiredTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
