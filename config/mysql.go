package config

import (
	"fmt"

	users "construct-week1/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "mysql911",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "db_construct1",
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
}
