package users

import "time"

type Core struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	Alamat       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//	business
type Business interface {

	//	Create user
	InsertData(data Core) (row int, err error)

	//	Get all user
	GetAllData() (data []Core, err error)

	//	Get user by ID
	GetDatabyID(id uint) (interface{}, error)
}

//	data
type Data interface {

	//	Create user
	InsertData(data Core) (row int, err error)

	//	Get all user
	SelectData() (data []Core, err error)

	//	Get user by ID
	SelectDatabyID(id uint) (interface{}, error)
}
