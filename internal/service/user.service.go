package service

import (
	"fmt"
	"time"

	"pawtopia.com/global"
	"pawtopia.com/internal/repo"
	"pawtopia.com/internal/ultils/crypto"
	"pawtopia.com/internal/ultils/random"
	"pawtopia.com/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	hashEmail := crypto.Encrypt(email)
	if hashEmail == "" {
		global.Logger.Error("Hash email error")
		return response.ErrorCodeHashEmailError
	}

	fmt.Printf("Hash email: %s\n", hashEmail)

	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeEmailExist
	}

	otp := random.GenerateSixLetterOTP()
	redisOTPErr := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if redisOTPErr != nil {
		global.Logger.Error(fmt.Sprintf("Redis error: %v", redisOTPErr))
		return response.ErrorCodeOTPError
	}
	return response.ErrorCodeSuccess
}
