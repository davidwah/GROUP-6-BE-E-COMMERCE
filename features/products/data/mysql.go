package data

import (
	"fmt"

	"construct-week1/features/products"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// InsertProduct implements products.Data
func (repo *mysqlProductRepository) InsertProduct(input products.Core) (int, error) {
	product := fromCore(input)

	res := repo.db.Create(&product)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, fmt.Errorf("Failed to insert product")
	}

	return int(res.RowsAffected), nil

}

// SelectProductbyID implements products.Data
func (repo *mysqlProductRepository) SelectProductbyID(id uint) ([]products.Core, error) {
	panic("unimplemented")
}

// Dependency Injection
func NewProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}
