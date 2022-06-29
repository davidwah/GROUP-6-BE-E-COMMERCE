package data

import (
	"errors"

	"construct-week1/features/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// InsertData implements users.Data
func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := fromCore(input)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return 0, errors.New("failed")
	}
	user.Password = string(bytes)

	res := repo.db.Create(&user)	// Check register belum ketemu
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, errors.New("failed")
	}

	return int(res.RowsAffected), nil
}

// SelectDatabyID implements users.Data
func (repo *mysqlUserRepository) SelectDatabyID(id int) (interface{}, error) {
	var userId User
	res := repo.db.First(&userId, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return userId, nil
}

// Dependency Injection
func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}
