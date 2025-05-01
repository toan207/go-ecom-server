package repo

import (
	"pawtopia.com/global"
	"pawtopia.com/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	row := global.MySQL.Table(TableNameUser).Select("username").Where("username = ?", email).First(&model.GoDbUser{}).Row()
	return row != nil
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
