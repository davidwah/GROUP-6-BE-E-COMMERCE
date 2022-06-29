package routes

import (
	"construct-week1/factory"
	"construct-week1/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	//	Login
	e.POST("/login", presenter.LoginPresenter.Login)

	//	User
	e.POST("/users", presenter.UserPresenter.Register)
	e.GET("/users/:id", presenter.UserPresenter.GetUserID, middlewares.JWTMiddleware())

	//	Product
	e.POST("/products", presenter.ProductPresenter.AddProduct, middlewares.JWTMiddleware())
	e.GET("/products", presenter.ProductPresenter.GetAllProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", presenter.ProductPresenter.GetProductbyID, middlewares.JWTMiddleware())
	
	return e
}
