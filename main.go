package main

import (
	"construct-week1/config"
	"construct-week1/factory"
	"construct-week1/routes"
	"fmt"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)
	fmt.Println("Server jalan di -> http://localhost:8081")
	e.Logger.Fatal(e.Start(":8081"))
}
