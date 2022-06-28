package main

import (
	"construct-week1/config"
	"construct-week1/factory"
	"construct-week1/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Logger.Fatal(e.Start(":80"))
}
