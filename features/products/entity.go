package products

import "time"

type Core struct {
	ID          int
	Name        string
	Price       int
	Quantity    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Bussiness interface {
	InsertData(data Core) (row int, err error)
	GetAllData() (data []Core, err error)
	GetDatabyId() (data []Core, err error)
}
