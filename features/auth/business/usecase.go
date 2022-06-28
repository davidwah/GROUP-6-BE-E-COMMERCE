package business

import "construct-week1/features/auth"

type authUseCase struct {
	authData auth.Data
}

// LoginUsers implements auth.Business
func (uc *authUseCase) LoginUsers(data auth.User, input string) (interface{}, error) {
	res, err := uc.authData.SelectLogin(data, input)
	return res, err
}

// Dependency Injection
func NewAuthBusiness(athData auth.Data) auth.Business {
	return &authUseCase{
		authData: athData,
	}
}
