package data

import (
	"construct-week1/features/products"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name        string    `json:"name" from:"name"`
	Price       int       `json:"price" from:"price"`
	Quantity    int       `json:"quantity" from:"quantity"`
	Description string    `json:"description" from:"description"`
	CreatedAt   time.Time `json:"created_at" from:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" from:"updated_at"`
}

// Data transfer object
func (data *Product) toCore() products.Core {
	return products.Core{
		ID:          int(data.ID),
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

// select data Products
func toCoreList(data []Product) []products.Core {
	res := []products.Core{}
	for key := range data {
		res = append(res, data[key].toCore())
	}
	return res
}

// insert Data Products
func fromCore(core products.Core) Product {
	return Product{
		Name:        core.Name,
		Price:       core.Price,
		Quantity:    core.Quantity,
		Description: core.Description,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}
