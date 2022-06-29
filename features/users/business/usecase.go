package business

import (
	"construct-week1/features/users"
	"errors"
)

type userUseCase struct {
	userData users.Data
}

// GetDatabyID implements users.Business
func (uc *userUseCase) GetDatabyID(id int) (interface{}, error) {
	response, err := uc.userData.SelectDatabyID(id)
	return response, err
}

// InsertData implements users.Business
func (uc *userUseCase) InsertData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Email == "" || input.Password == "" || input.Alamat == "" {
		return -1, errors.New("all input data must be filled")
	}

	row, err = uc.userData.InsertData(input)
	return row, err
}

// Dependency Injection
func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}
