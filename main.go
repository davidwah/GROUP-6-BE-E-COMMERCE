package main

import (
	"construct-week1/config"
	"construct-week1/factory"
	"construct-week1/routes"
	"construct-week1/middleware"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)
	middleware.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":80"))
}
