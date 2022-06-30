package orders

type Core struct {
	ID         uint
	UserId     uint
	Address    Address
	CreditCart CreditCart
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    uint
}

type CreditCart struct {
	Type   string
	Name   string
	Number string
	CVV    uint
	Month  uint
	Year   uint
}

type Business interface{}

type Data interface{}