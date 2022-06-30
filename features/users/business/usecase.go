package business

import (
	"construct-week1/features/users"
	"errors"
)

type userUseCase struct {
	userData users.Data
}

// UpdatedData implements users.Business
func (uc *userUseCase) UpdatedData(id int, input users.Core) (error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("all input data must be filled")
	}

	row := uc.userData.UpdateData(id, input)
	return row
}

// GetDatabyID implements users.Business
func (uc *userUseCase) FindData(id int) (interface{}, error) {
	response, err := uc.userData.SelectData(id)
	return response, err
}

// InsertData implements users.Business
func (uc *userUseCase) InsertData(input users.Core) (int, error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return -1, errors.New("all input data must be filled")
	}

	row, err := uc.userData.InsertData(input)
	return row, err
}

// Dependency Injection
func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}
