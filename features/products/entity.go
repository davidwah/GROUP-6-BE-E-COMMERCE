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

	UpdateProduct(id string, data map[string]interface{}) (error)

	DeleteProduct(id int) error

	GetProduct(limit, offset int) ([]Core, error)

	GetProductbyID(id uint) ([]Core, error)

	GetProductID (id int) (interface{}, error)
}

//	data
type Data interface {

	InsertProductData(Core) (int, error)

	UpdateProductData(id string, data map[string]interface{}) (error)

	DeleteProductData(id int) error

	SelectProductData() ([]Core, error)

	SelectProductbyIDData(id uint) ([]Core, error)

	SelectProductID (id int) (interface{}, error)
}
