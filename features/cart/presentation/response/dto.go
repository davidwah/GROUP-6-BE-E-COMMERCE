package response

import (
	"construct-week1/features/cart"
)

type Cart struct {
	ID      uint    `json:"id"`
	Qty     uint    `json:"qty"`
	UserId  uint    `json:"user_id"`
	Product Product `json:"product"`
}

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
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
