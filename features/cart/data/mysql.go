package data

import (
	"construct-week1/features/cart"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

// SelectCart implements cart.Data
func (repo *mysqlCartRepository) SelectCart(id uint) ([]cart.Core, error) {
	var dataCart []Cart
	result := repo.db.Preload("Product").Where("user_id = ?", id).Find(&dataCart)
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

func (repo *mysqlCartRepository) UpdateCartData(idCart uint, idUser uint, data cart.CoreUpdate) error {
	
	res := repo.db.Table("carts").Where("id = ? AND user_id = ?", idCart, idUser).Updates(data)
	fmt.Println(res)
	if res.Error == nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.New("Failed")
	}

	return nil
}

// Dependency Injection
func NewCartRepository(conn *gorm.DB) cart.Data {
	return &mysqlCartRepository{
		db: conn,
	}
}
