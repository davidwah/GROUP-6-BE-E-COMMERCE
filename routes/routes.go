package routes

import (
	"construct-week1/factory"
	"construct-week1/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	//	Login
	e.POST("/login", presenter.LoginPresenter.Login)

	//	User
	e.POST("/users", presenter.UserPresenter.Register)
	e.GET("/users/:id", presenter.UserPresenter.GetUserID, middlewares.JWTMiddleware())

	//	Product
	e.POST("/products", presenter.ProductPresenter.AddProduct, middlewares.JWTMiddleware())
	e.GET("/products", presenter.ProductPresenter.GetAllProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", presenter.ProductPresenter.GetProductbyID, middlewares.JWTMiddleware())

	//	Cart
	e.POST("/cart", presenter.CartPresenter.AddCart, middlewares.JWTMiddleware())
	e.GET("/cart", presenter.CartPresenter.FindCart, middlewares.JWTMiddleware())

	return e
}
