package data

import (
	"construct-week1/features/cart"
	"fmt"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

// SelectCart implements cart.Data
func (repo *mysqlCartRepository) SelectCart(id int) ([]cart.Core, error) {
	var dataCart []Cart
	result := repo.db.Preload("Product").Find(&dataCart)
	if result.Error != nil {
		return []cart.Core{}, result.Error
	}

	return toCoreList(dataCart), nil
}

// InsertCartData implements cart.Data
func (repo *mysqlCartRepository) InsertCartData(data cart.Core) error {
	//	Kondisi ketika quantity yang diinginkan melebihi ketersediaan
	cart := fromCore(data)

	res := repo.db.Create(&cart)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return fmt.Errorf("failed to insert cart")
	}

	return nil
}

// Dependency Injection
func NewCartRepository(conn *gorm.DB) cart.Data {
	return &mysqlCartRepository{
		db: conn,
	}
}
