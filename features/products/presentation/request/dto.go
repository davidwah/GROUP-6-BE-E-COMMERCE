package request

import "construct-week1/features/products"

type Product struct {
	// ID          uint   `json:"id" form:"id"`
	IDUser      uint   `json:"user_id" form:"user_id"`
	Name        string `json:"name" form:"name"`
	Price       uint   `json:"price" form:"price"`
	Quantity    uint   `json:"qty" form:"qty"`
	Description string `json:"desc" form:"desc"`
}

func ToCore(req Product) products.Core {
	return products.Core{
		IDUser:      req.IDUser,
		Name:        req.Name,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Description: req.Description,
	}
}

func DisplayAdd(req Product) products.Core {
	return products.Core{
		Name:        req.Name,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Description: req.Description,
	}
}
