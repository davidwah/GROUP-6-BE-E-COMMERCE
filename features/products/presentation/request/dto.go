package request

import "construct-week1/features/products"

type Product struct {
	IDUser      uint   `json:"id_user" form:"id_user"`
	Name        string `json:"name" form:"name"`
	Price       uint   `json:"price" form:"price"`
	Quantity    uint   `json:"qty" form:"qty"`
	Description string `json:"desc" form:"desc"`
}

// type ProductUpdate struct {
// 	ID          uint   `json:"id" form:"id"`
// 	IDUser      uint   `json:"user_id" form:"user_id"`
// 	Name        string `json:"name" form:"name"`
// 	Price       uint   `json:"price" form:"price"`
// 	Quantity    uint   `json:"qty" form:"qty"`
// 	Description string `json:"desc" form:"desc"`
// }

func ToCore(req Product) products.Core {
	return products.Core{
		IDUser:      req.IDUser,
		Name:        req.Name,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Description: req.Description,
	}
}

// func ToCoreUpdate(req ProductUpdate) products.Core {
// 	return products.Core{
// 		ID:          req.ID,
// 		IDUser:      req.IDUser,
// 		Name:        req.Name,
// 		Price:       req.Price,
// 		Quantity:    req.Quantity,
// 		Description: req.Description,
// 	}
// }
