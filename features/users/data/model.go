package data

import (
	"construct-week1/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Alamat       string `json:"alamat" form:"alamat"`
}

//	DTO (Data Transfer Object)
func (data *User) toCore() users.Core {
	return users.Core{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		Password:     data.Password,
		Alamat:       data.Alamat,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

//	Digunakan dalam pembuatan SelectData (mysql.go)
func toCoreList(data []User) []users.Core {
	res := []users.Core{}
	for key := range data {
		res = append(res, data[key].toCore())
	}
	return res
}

//	Digunakan dalam pembuatan InsertData (mysql.go)
func fromCore(core users.Core) User {
	return User{
		Name:         core.Name,
		Email:        core.Email,
		Password:     core.Password,
		Alamat:       core.Alamat,
	}
}
