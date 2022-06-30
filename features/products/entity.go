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
	InsertProduct(Core) (int, error)

	UpdateProduct(id string, data map[string]interface{}) error

	DeleteProduct(id int) error

	FindAllProduct(limit, offset int) ([]Core, error)

	FindProductbyIDUser(id uint) ([]Core, error)

	FindProductbyIDProduct(id uint) (interface{}, error)
}

//	data
type Data interface {
	InsertProductData(Core) (int, error)

	UpdateProductData(id string, data map[string]interface{}) error

	DeleteProductData(id int) error

	SelectProduct() ([]Core, error)

	SelectProductbyIDUser(id uint) ([]Core, error)

	SelectProductbyIDProduct(id uint) (interface{}, error)
}
