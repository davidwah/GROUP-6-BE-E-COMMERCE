package business

import (
	"construct-week1/features/users"
	"errors"
)

type userUseCase struct {
	userData users.Data
}

// GetAllData implements users.Business
func (uc *userUseCase) GetAllData() (data []users.Core, err error) {
	response, err := uc.userData.SelectData()
	return response, err
}

// GetDatabyID implements users.Business
func (uc *userUseCase) GetDatabyID(id uint) (interface{}, error) {
	response, err := uc.userData.SelectDatabyID(id)
	return response, err
}

// InsertData implements users.Business
func (uc *userUseCase) InsertData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Username == "" || input.Email == "" || input.Password == "" || input.Alamat == "" || input.NoHandPhone == "" || input.JenisKelamin == "" || input.TanggalLahir == "" {
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
