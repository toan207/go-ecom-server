package po

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `gorm:"column:uuid;type:char(36);primary_key" json:"uuid"`
	Username string    `gorm:"column:username;type:varchar(50);unique;not null" json:"username"`
	Password string    `gorm:"column:password;type:varchar(100);not null" json:"password"`
	IsActive bool      `gorm:"column:is_active;type:tinyint(1);default:1" json:"is_active"`
	Roles    []Role    `gorm:"many2many:go_db_user_roles" json:"roles"`
}

func (u *User) TableName() string {
	return "go_db_users"
}
