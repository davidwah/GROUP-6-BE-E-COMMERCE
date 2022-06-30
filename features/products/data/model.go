package data

import (
	"construct-week1/features/products"
	// "construct-week1/features/users"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
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

//	(Mapping) gorm.Model (User struct) to Core (entity.go)
func (data *Product) toCore() products.Core {
	return products.Core{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
	}
}

//	Menampilkan detail data book dari beberapa book
//	Digunakan dalam pembuatan SelectData (mysql.go)
func toCoreList(data []Product) []products.Core {
	result := []products.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

//	(Mapping) Core (entity.go) to gorm.Model (User struct)
//	Digunakan dalam pembuatan InsertData (mysql.go)
func fromCore(core products.Core) Product {
	return Product{
		IDUser:      core.IDUser,
		Name:        core.Name,
		Price:       core.Price,
		Quantity:    core.Quantity,
		Description: core.Description,
	}
}

// func fromCoreUpdate(core products.Core) ProductUpdate {
// 	return ProductUpdate{
// 		ID:          core.ID,
// 		IDUser:      core.IDUser,
// 		Name:        core.Name,
// 		Price:       core.Price,
// 		Quantity:    core.Quantity,
// 		Description: core.Description,
// 	}
// }
