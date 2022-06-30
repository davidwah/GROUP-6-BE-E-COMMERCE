package config

import (
	"fmt"

	carts "construct-week1/features/cart/data"
	products "construct-week1/features/products/data"
	users "construct-week1/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	db_username := "root"
	db_password := "mysql911"
	db_port := "3306"
	db_host := "127.0.0.1"
	db_name := "db_construct1"

	config := map[string]interface{}{

		"DB_Username": db_username,
		"DB_Password": db_password,
		"DB_Port":     db_port,
		"DB_Host":     db_host,
		"DB_Name":     db_name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate(db)

	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&products.Product{})
	db.AutoMigrate(&carts.Cart{})
}
