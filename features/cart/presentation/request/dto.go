package request

import (
	"construct-week1/features/cart"
)

type Cart struct {
	ID        uint `json:"id"`
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	Qty       uint `json:"qty"`
	Status    int  `json:"status"`
}

func ToCore(req Cart) cart.Core {
	return cart.Core{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Qty:       req.Qty,
		Status:    req.Status,
	}
}
