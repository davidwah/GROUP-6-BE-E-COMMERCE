package business

import "construct-week1/features/products"

type productUseCase struct {
	productData products.Data
}

// GetProductbyID implements products.Business
func (uc *productUseCase) GetProductbyID(id uint) ([]products.Core, error) {
	panic("unimplemented")
}

// InsertProduct implements products.Business
func (uc *productUseCase) InsertProduct(data products.Core) (int, error) {
	panic("unimplemented")
}

// Dependency Injection
func NewProductBusiness(pdctData products.Data) products.Business {
	return &productUseCase{
		productData: pdctData,
	}
}
