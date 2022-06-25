package data

import (
	"construct-week1/features/users"
	"errors"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// InsertData implements users.Data
func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := fromCore(input)

	res := repo.db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, errors.New("failed")
	}

	return int(res.RowsAffected), nil
}

// SelectData implements users.Data
func (repo *mysqlUserRepository) SelectData() (data []users.Core, err error) {
	dataUsers := []User{}
	res := repo.db.Find(&dataUsers)
	if res.Error != nil {
		return []users.Core{}, errors.New("failed")
	}

	return toCoreList(dataUsers), nil
}

// SelectDatabyID implements users.Data
func (repo *mysqlUserRepository) SelectDatabyID(id uint) (interface{}, error) {
	userId := User{}
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
