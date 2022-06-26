package business

import "construct-week1/features/auth"

type loginUseCase struct {
	authData auth.Data
}

// LoginUsers implements auth.Business
func (uc *loginUseCase) LoginUsers(data auth.User) (interface{}, error) {
	res, err := uc.authData.SelectLogin(data)
	return res, err
}

// Dependency Injection
func NewUserBusiness(athData auth.Data) auth.Business {
	return &loginUseCase{
		authData: athData,
	}
}
