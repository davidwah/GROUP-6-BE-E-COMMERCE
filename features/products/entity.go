package products

type Core struct {
	ID          uint
	IDUser      uint
	Name        string
	Price       uint
	Quantity    uint
	Description string
}

//	business
type Business interface {
	InsertProduct(data Core) (int, error)

	UpdateProduct(id string, data map[string]interface{}) error

	DeleteProduct(id int) error

	GetProduct(limit, offset int) ([]Core, error)

	GetProductbyID(id uint) ([]Core, error)
}

//	data
type Data interface {
	InsertProductData(data Core) (int, error)

	UpdateProductData(id string, data map[string]interface{}) error

	DeleteProductData(id int) error

	SelectProductData() ([]Core, error)

	SelectProductbyIDData(id uint) ([]Core, error)
}
