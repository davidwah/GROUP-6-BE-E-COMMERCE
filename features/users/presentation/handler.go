package presentation

import (
	"net/http"
	"strconv"

	"construct-week1/features/helper"
	"construct-week1/features/users"
	"construct-week1/features/users/presentation/request"
	"construct-week1/features/users/presentation/response"
	"construct-week1/middlewares"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

//	Menampilkan semua data user
//	Done
func (handle *UserHandler) GetAll(c echo.Context) error {

	res, err := handle.userBusiness.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get all data user"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get all user", response.FromCoreList(res)))

}

//	Menampilkan data user berdasarkan ID
//	Done
func (handle *UserHandler) GetUserID(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unathorized"))
	}

	userId, err := handle.userBusiness.GetDatabyID(uint(idUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed get your data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get your data", userId))
}

//	Belum selesai (kurang lengkap pada bagian extract email menggunakan regexp atau net/mail)
func (handle *UserHandler) Register(c echo.Context) error {
	var newuser request.User
	err := c.Bind(&newuser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to bind data user, check your input"))
	}
	
	dataUser := request.ToCore(newuser)
	row, err := handle.userBusiness.InsertData(dataUser)
	if row == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Email anda telah terdaftar"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to insert your data", row))
}
