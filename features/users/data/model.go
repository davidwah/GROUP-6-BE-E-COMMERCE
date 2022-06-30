package data

import (
	"construct-week1/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
}

//	Digunakan dalam pembuatan InsertData (mysql.go)
func fromCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}
