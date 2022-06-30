package business

import (
	"construct-week1/features/cart"
	"errors"
)

type cartUseCase struct {
	cartData cart.Data
}

// UpdateCart implements cart.Business
func (uc *cartUseCase) UpdateCart(idCart uint, idUser uint, data cart.CoreUpdate) error {
	if data.Qty == 0 {
		return errors.New("all input data must be filled")
	}

	row := uc.cartData.UpdateCartData(idCart, idUser, data)
	return row
}

// FindDataCart implements cart.Business
func (uc *cartUseCase) FindCart(id uint) ([]cart.Core, error) {
	resp, err := uc.cartData.SelectCart(id)
	return resp, err
}

// InsertCart implements cart.Business
func (uc *cartUseCase) InsertCart(data cart.Core) error {
	if data.ProductId == 0 || data.Qty == 0 {
		return errors.New("all input data must be filled")
	}

	row := uc.cartData.InsertCartData(data)
	return row
}

// Dependency Injection
func NewProductBusiness(cartData cart.Data) cart.Business {
	return &cartUseCase{
		cartData: cartData,
	}
}
