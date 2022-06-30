package users

// "construct-week1/features/cart"
// "construct-week1/features/products"

type Core struct {
	Name      string
	Email     string
	Password  string
}

//	business
type Business interface {

	//	Create user
	InsertData(Core) (int, error)

	//	Get user by ID
	FindData(id int) (interface{}, error)

	//	Updated user
	UpdatedData(id int, data Core) (error)
}

//	data
type Data interface {

	//	Create user
	InsertData(Core) (int, error)

	//	Get user by ID
	SelectData(id int) (interface{}, error)

	//	Updated user
	UpdateData(id int, data Core) (error)
}
