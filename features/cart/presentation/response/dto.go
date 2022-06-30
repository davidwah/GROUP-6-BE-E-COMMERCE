package response

import (
	"construct-week1/features/cart"
)

type Cart struct {
	ID      uint    `json:"id" form:"id"`
	Qty     uint    `json:"qty" form:"qty"`
	UserId  uint    `json:"id_user" form:"id_user"`
	Product Product `json:"product"`
}

type Product struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Price       uint   `json:"price" form:"price"`
	Quantity    uint   `json:"qty" form:"qty"`
	Description string `json:"desc" form:"desc"`
}

func FromCore(data cart.Core) Cart {
	return Cart{
		ID:     data.ID,
		Qty:    data.Qty,
		UserId: data.UserId,
		Product: Product{
			ID:          data.Product.ID,
			Name:        data.Product.Name,
			Price:       data.Product.Price,
			Quantity:    data.Product.Quantity,
			Description: data.Product.Description,
		},
	}
}

func FromCoreList(data []cart.Core) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
