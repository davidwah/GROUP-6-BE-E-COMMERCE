package products

import (
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

func (data *Product) toCore() products.Core {
	return products.Core{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func toCoreList(data []Product) []products.Core {
	res := []products.Core{}
	for key := range data {
		res = append(res, data[key].toCore())
	}
	return res
}
