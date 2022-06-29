package users

import (
	// "construct-week1/features/cart"
	// "construct-week1/features/products"
	"time"
)

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Products  []products.Core
	// Carts     []cart.Core
}

//	business
type Business interface {

	//	Create user
	InsertData(data Core) (row int, err error)

	//	Get user by ID
	GetDatabyID(id int) (interface{}, error)
}

//	data
type Data interface {

	//	Create user
	InsertData(data Core) (row int, err error)

	//	Get user by ID
	SelectDatabyID(id int) (interface{}, error)
}
