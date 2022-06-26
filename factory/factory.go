package factory

import (
	userBusiness "construct-week1/features/users/business"
	userData "construct-week1/features/users/data"
	userPresentation "construct-week1/features/users/presentation"

	loginBusiness "construct-week1/features/auth/business"
	loginData "construct-week1/features/auth/data"
	loginPresentation "construct-week1/features/auth/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *userPresentation.UserHandler
	LoginPresenter *loginPresentation.AuthHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	loginData := loginData.NewUserRepository(dbConn)
	loginBusiness := loginBusiness.NewUserBusiness(loginData)
	loginPresentation := loginPresentation.NewUserHandler(loginBusiness)

	userData := userData.NewUserRepository(dbConn)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		LoginPresenter: loginPresentation,
		UserPresenter: userPresentation,
	}
}

