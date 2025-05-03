package repo

import (
	"pawtopia.com/global"
	database "pawtopia.com/internal/databse"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	// row := global.MySQL.Table(TableNameUser).Select("username").Where("username = ?", email).First(&model.GoDbUser{}).Row()
	user, err := u.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.MySQLC),
	}
}
