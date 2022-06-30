package request

import (
	"construct-week1/features/cart"
)

type Cart struct {
	ID        uint `json:"id" form:"id"`
	ProductId uint `json:"id_product" form:"id_product"`
	UserId    uint `json:"id_user" form:"id_user"`
	Qty       uint `json:"qty" form:"qty"`
	Status    int  `json:"status" form:"status"`
}

func ToCore(req Cart) cart.Core {
	return cart.Core{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Qty:       req.Qty,
		Status:    req.Status,
	}
}
