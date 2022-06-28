package routes

import (
	"construct-week1/factory"

	"github.com/labstack/echo"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	//	Login
	e.POST("/login", presenter.LoginPresenter.Login)

	//	User
	e.POST("/users", presenter.UserPresenter.Register)
	e.GET("/users", presenter.UserPresenter.GetAll)
	e.GET("/user/:id", presenter.UserPresenter.GetUserID)

	//	Prodcut
	e.POST("/products", presenter.ProductPresenter.AddProduct)


	return e
}