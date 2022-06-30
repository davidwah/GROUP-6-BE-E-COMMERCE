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
	e.GET("/users/:id", presenter.UserPresenter.FindByID, middlewares.JWTMiddleware())
	e.PUT("/users/:id", presenter.UserPresenter.Updated, middlewares.JWTMiddleware())

	//	Product
	e.POST("/products", presenter.ProductPresenter.AddProduct, middlewares.JWTMiddleware())
	e.GET("/products", presenter.ProductPresenter.FindProduct)
	e.GET("/products/:id", presenter.ProductPresenter.FindProductbyID, middlewares.JWTMiddleware())
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateProduct, middlewares.JWTMiddleware())

	//	Cart
	e.POST("/cart", presenter.CartPresenter.AddCart, middlewares.JWTMiddleware())
	e.GET("/cart", presenter.CartPresenter.FindCart, middlewares.JWTMiddleware())

	return e
}
