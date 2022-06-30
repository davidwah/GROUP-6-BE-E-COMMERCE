package data

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId     uint       `json:"id_user" form:"id_user"`
	Address    Address    `json:"address" form:"address"`
	CreditCart CreditCart `json:"credit_cart" form:"credit_cart"`
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
