// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGoDbUser = "go_db_users"

// GoDbUser mapped from table <go_db_users>
type GoDbUser struct {
	UUID     string `gorm:"column:uuid;primaryKey" json:"uuid"`
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
	IsActive bool   `gorm:"column:is_active;default:1" json:"is_active"`
}

// TableName GoDbUser's table name
func (*GoDbUser) TableName() string {
	return TableNameGoDbUser
}
