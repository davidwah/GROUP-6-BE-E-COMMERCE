package data

import (
	"fmt"

	"construct-week1/features/products"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// DeleteProductData implements products.Data
func (*mysqlProductRepository) DeleteProductData(id int) error {
	panic("unimplemented")
}

// SelectProductData implements products.Data
func (repo *mysqlProductRepository) SelectProductData() ([]products.Core, error) {
	var dataProducts []Product

	res := repo.db.Find(&dataProducts)
	if res.Error != nil {
		return []products.Core{}, res.Error
	}
	return toCoreList(dataProducts), nil
}

// UpdateProductData implements products.Data
func (*mysqlProductRepository) UpdateProductData(id string, data map[string]interface{}) error {
	panic("unimplemented")
}

// InsertProduct implements products.Data
func (repo *mysqlProductRepository) InsertProductData(input products.Core) (int, error) {
	product := fromCore(input)

	res := repo.db.Create(&product)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert product")
	}

	return int(res.RowsAffected), nil
}

// SelectProductbyID implements products.Data
func (repo *mysqlProductRepository) SelectProductbyIDData(id uint) ([]products.Core, error) {
	var productsId []Product
	res := repo.db.Table("products").Select("*").Where("id_user", id).Find(&productsId)
	if res.Error != nil {
		return []products.Core{}, res.Error
	}

	return toCoreList(productsId), nil
}

// Dependency Injection
func NewProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}
