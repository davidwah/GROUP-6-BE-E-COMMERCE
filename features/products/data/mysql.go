package data

import (
	"construct-week1/features/products"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// Insert data Product
func (repo *mysqlProductRepository) InsertData(input products.Core) (row int, err error) {
	product := fromCore(input)

	res := repo.db.Create(&product)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, errors.New("")
	}
	fmt.Println(res.RowsAffected)
	return int(res.RowsAffected), nil
}

// Select data Products
func (repo *mysqlProductRepository) SelectData() (data []products.Core, err error) {
	dataProducts := []Product{}
	res := repo.db.Find(&dataProducts)
	if res.Error != nil {
		return []products.Core{}, errors.New("gagal")
	}
	return toCoreList(dataProducts), nil
}

// select data by Id
func (repo *mysqlProductRepository) selecDataById(id uint) (interface{}, error) {
	productId := Product{}
	res := repo.db.First(&productId, id)
	if res.Error != nil {
		return res.Error, nil
	}
	return productId, nil
}

// Dependency
func NewProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}
