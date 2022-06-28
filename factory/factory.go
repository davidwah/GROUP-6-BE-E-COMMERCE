package factory

import (
	userBusiness "construct-week1/features/users/business"
	userData "construct-week1/features/users/data"
	userPresentation "construct-week1/features/users/presentation"

	productBusiness "construct-week1/features/products/business"
	productData "construct-week1/features/products/data"
	productPresentation "construct-week1/features/products/presentation"

	loginBusiness "construct-week1/features/auth/business"
	loginData "construct-week1/features/auth/data"
	loginPresentation "construct-week1/features/auth/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	ProductPresenter *productPresentation.ProductHandler
	UserPresenter    *userPresentation.UserHandler
	LoginPresenter   *loginPresentation.AuthHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	loginData := loginData.NewAuthRepository(dbConn)
	loginBusiness := loginBusiness.NewAuthBusiness(loginData)
	loginPresentation := loginPresentation.NewAuthHandler(loginBusiness)

	userData := userData.NewUserRepository(dbConn)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	productData := productData.NewProductRepository(dbConn)
	productBusiness := productBusiness.NewProductBusiness(productData)
	productPresentation := productPresentation.NewProductHandler(productBusiness)

	return Presenter{
		LoginPresenter:   loginPresentation,
		UserPresenter:    userPresentation,
		ProductPresenter: productPresentation,
	}
}
