package data

import (
	"errors"

	"construct-week1/features/auth"
	"construct-week1/middlewares"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

// SelectLogin implements auth.Data
func (repo *mysqlAuthRepository) SelectLogin(data auth.User, input string) (interface{}, error) {
	// SelectLogin implements auth.Data
	var userData auth.User
	err := repo.db.Where("email = ?", data.Email).First(&userData)
	if err.Error != nil {
		return nil, err.Error
	}

	res := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input))
	if res != nil {
		return nil, errors.New("failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return nil, errToken
	}

	return token, nil

}

// Dependency Injection
func NewAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}
