package request

import (
	"construct-week1/features/users"
	"time"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Alamat       string `json:"alamat" form:"alamat"`
}

// Digunakan dalam pembuatan AddUser (handler.go) yang ditampilkan di Postman
func ToCore(req User) users.Core {
	return users.Core{
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		Alamat:       req.Alamat,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
