package response

import (
	"construct-week1/features/users"
	"time"
)

type User struct {
	ID           uint      `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	JenisKelamin string    `json:"jeniskelamin" form:"jeniskelamin"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

// Digunakan dalam pembuatan GetUserID (handler.go)
func FromCore(data users.Core) User {
	return User{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		JenisKelamin: data.JenisKelamin,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

// Digunakan dalam pembuatan GetAll (handler.go)
func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
