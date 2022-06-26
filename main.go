package main

import (
	"construct-week1/config"
	"construct-week1/factory"
	"construct-week1/routes"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)
	e.Logger.Fatal(e.Start(":80"))
}