package data

import (
	"construct-week1/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Username     string `json:"username" form:"username"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Alamat       string `json:"alamat" form:"alamat"`
	NoHandPhone  string `json:"nohp" form:"nohp"`
	JenisKelamin string `json:"jeniskelamin" form:"jeniskelamin"`
	TanggalLahir string `json:"tanggallahir" form:"tanggallahir"`
}

//	DTO (Data Transfer Object)
func (data *User) toCore() users.Core {
	return users.Core{
		ID:           data.ID,
		Name:         data.Name,
		Username:     data.Username,
		Email:        data.Email,
		Password:     data.Password,
		Alamat:       data.Alamat,
		NoHandPhone:  data.NoHandPhone,
		JenisKelamin: data.JenisKelamin,
		TanggalLahir: data.TanggalLahir,
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
		Username:     core.Username,
		Email:        core.Email,
		Password:     core.Password,
		Alamat:       core.Alamat,
		NoHandPhone:  core.NoHandPhone,
		JenisKelamin: core.JenisKelamin,
		TanggalLahir: core.TanggalLahir,
	}
}
