package users

import "time"

/*
	Data diri User
	1. Name
	2. UserName
	3. Email
	4. Password
	5. Alamat
	6. No HandPhone
	7. Jenis Kelamin
	8. Tanggal Lahir
*/

type Core struct {
	ID           uint
	Name         string
	Username     string
	Email        string
	Password     string
	Alamat       string
	NoHandPhone  string
	JenisKelamin string
	TanggalLahir string
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
