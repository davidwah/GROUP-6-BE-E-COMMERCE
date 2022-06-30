package business

import (
	"construct-week1/features/products"
	"errors"
)

type productUseCase struct {
	productData products.Data
}

// GetProductID implements products.Business
func (uc *productUseCase) GetProductID(id int) (interface{}, error) {
	resp, err := uc.productData.SelectProductID(id)
	return resp, err
}

// DeleteProduct implements products.Business
func (uc *productUseCase) DeleteProduct(id int) error {
	req := uc.productData.DeleteProductData(id)
	return req
}

// GetProduct implements products.Business
func (uc *productUseCase) GetProduct(limit, offset int) ([]products.Core, error) {
	resp, err := uc.productData.SelectProductData()
	return resp, err
}

// GetProductbyID implements products.Business
func (uc *productUseCase) GetProductbyID(id uint) ([]products.Core, error) {
	resp, err := uc.productData.SelectProductbyIDData(id)
	return resp, err
}

// InsertProduct implements products.Business
func (uc *productUseCase) InsertProduct(data products.Core) (int, error) {
	if data.Name == "" || data.Price == 0 || data.Quantity == 0 || data.Description == "" {
		return -1, errors.New("all input data must be filled")
	}

	row, err := uc.productData.InsertProductData(data)
	return row, err
}

// UpdateProduct implements products.Business
func (uc *productUseCase) UpdateProduct(id string, data map[string]interface{}) error {
	req := uc.productData.UpdateProductData(id, data)
	return req
}

// Dependency Injection
func NewProductBusiness(pdctData products.Data) products.Business {
	return &productUseCase{
		productData: pdctData,
	}
}
