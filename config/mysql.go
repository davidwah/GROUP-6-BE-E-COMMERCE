package config

import (
	"fmt"
	"os"

	products "construct-week1/features/products/data"
	users "construct-week1/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": os.Getenv("username"),
		"DB_Password": os.Getenv("password"),
		"DB_Port":     os.Getenv("port"),
		"DB_Host":     os.Getenv("host"),
		"DB_Name":     os.Getenv("db_name"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

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
}
