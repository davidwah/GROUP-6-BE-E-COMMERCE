package data

import (
	"construct-week1/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	Qty       uint `json:"qty"`
	Status    int  `json:"status"`
	Product   Product
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"qty" form:"qty"`
	Description string `json:"desc" form:"desc"`
	Carts       []Cart
}

func (data *Cart) toCore() cart.Core {
	return cart.Core{
		ProductId: data.ProductId,
		UserId:    data.UserId,
		Qty:       data.Qty,
		Product: cart.Product{
			ID:          data.Product.ID,
			Name:        data.Product.Name,
			Price:       data.Product.Price,
			Quantity:    data.Product.Quantity,
			Description: data.Product.Description,
		},
	}
}

func toCoreList(data []Cart) []cart.Core {
	result := []cart.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core cart.Core) Cart {
	return Cart{
		ProductId: core.ProductId,
		UserId:    core.UserId,
		Qty:       core.Qty,
		Status:    core.Status,
	}
}
