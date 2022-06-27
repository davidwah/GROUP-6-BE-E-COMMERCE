package products

import "time"

type Core struct {
	ID          uint
	Name        string
	Price       uint
	Quantity    uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//	business
type Business interface {

	InsertProduct(data Core) (int, error)

	GetProductbyID(id uint) ([]Core, error)
}

//	data
type Data interface {

	InsertProduct(data Core) (int, error)

	SelectProductbyID(id uint) ([]Core, error)
}
