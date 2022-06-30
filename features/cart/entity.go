package cart

type Core struct {
	ID        uint
	ProductId uint
	UserId    uint
	Qty       uint
	Status    int // 0 (paid) / 1 (unpaid)
	Product   Product
}

type CoreUpdate struct {
	Qty uint
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
	FindCart(id uint) ([]Core, error)

	//	Update cart by ID Cart
	UpdateCart(idCart uint, idUser uint, data CoreUpdate) error
}

type Data interface {
	//	Create cart
	InsertCartData(Core) error

	//	Get cart by ID
	SelectCart(id uint) ([]Core, error)

	//	Update cart by ID Cart
	UpdateCartData(idCart uint, idUser uint, data CoreUpdate) error
}
