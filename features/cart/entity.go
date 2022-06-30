package cart

type Core struct {
	ID        uint
	ProductId uint
	UserId    uint
	Qty       uint
	Status    int // 0 (paid) / 1 (unpaid)
	Product   Product
}

type Product struct {
	ID          uint
	Name        string
	Price       uint
	Quantity    uint
	Description string
}

type Business interface {
	//	Create cart
	InsertCart(Core) error

	//	Get cart by ID
	FindDataCart(id int) ([]Core, error)
}

type Data interface {
	//	Create cart
	InsertCartData(Core) error

	//	Get cart by ID
	SelectCart(id int) ([]Core, error)
}
