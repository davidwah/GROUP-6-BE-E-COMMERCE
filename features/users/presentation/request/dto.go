package request

import (
	"construct-week1/features/users"
	"time"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Username     string `json:"username" form:"username"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Alamat       string `json:"alamat" form:"alamat"`
	NoHandPhone  string `json:"nohp" form:"nohp"`
	JenisKelamin string `json:"jeniskelamin" form:"jeniskelamin"`
	TanggalLahir string `json:"tanggallahir" form:"tanggallahir"`
}

// Digunakan dalam pembuatan AddUser (handler.go) yang ditampilkan di Postman
func ToCore(req User) users.Core {
	return users.Core{
		Name:         req.Name,
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		Alamat:       req.Alamat,
		NoHandPhone:  req.NoHandPhone,
		JenisKelamin: req.JenisKelamin,
		TanggalLahir: req.TanggalLahir,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
