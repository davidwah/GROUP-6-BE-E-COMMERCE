package data

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId     uint `json:"id_user" form:"id_user"`
	Address    Address
	CreditCart CreditCart
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
	Zip    uint
}

type CreditCart struct {
	gorm.Model
	Type   string
	Name   string
	Number string
	CVV    uint
	Month  uint
	Year   uint
}
