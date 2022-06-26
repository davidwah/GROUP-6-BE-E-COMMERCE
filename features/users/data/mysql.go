package data

import (
	"construct-week1/features/users"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// CheckRegister implements users.Data
// func (repo *mysqlUserRepository) CheckRegister(dataCheck map[string]string) bool {
// 	checkUser := User{}
// 	res := repo.db.First(&checkUser, dataCheck)

// 	return res == nil
// }

func (repo *mysqlUserRepository) CheckRegister(dataCheck string) int {
	checkUser := User{}
	res := repo.db.Where("email = ?", dataCheck).FirstOrCreate(&checkUser)
	if res.RowsAffected != 0 {
		return 1
	}

	return 0
}

// InsertData implements users.Data
func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := fromCore(input)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return 0, errors.New("failed")
	}
	user.Password = string(bytes)
	
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
