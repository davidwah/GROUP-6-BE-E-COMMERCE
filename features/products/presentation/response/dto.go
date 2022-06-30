package response

import "construct-week1/features/products"

type Product struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Price       uint      `json:"price" form:"price"`
	Quantity    uint      `json:"qty" form:"qty"`
	Description string    `json:"desc" form:"desc"`
}

type ProductID struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Price       uint      `json:"price" form:"price"`
	Quantity    uint      `json:"qty" form:"qty"`
	Description string    `json:"desc" form:"desc"`
}

// (Mapping) Core to Product
func FromCore(data products.Core) Product {
	return Product{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
	}
}

// (Mapping) Core to Products
// Digunakan dalam pembuatan GetAll (handler.go)
func FromCoreList(data []products.Core) []Product {
	result := []Product{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCoreGetID(data products.Core) ProductID {
	return ProductID{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
	}
}

func FromCoreListGetID(data []products.Core) []ProductID {
	result := []ProductID{}
	for key := range data {
		result = append(result, FromCoreGetID(data[key]))
	}
	return result
}
