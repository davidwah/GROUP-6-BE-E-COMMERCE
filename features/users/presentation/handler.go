package presentation

import (
	"construct-week1/features/users"
	"construct-week1/features/users/presentation/response"
	"construct-week1/middlewares"
	"net/http"
	"strconv"

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
func (handle *UserHandler) GetAll(c echo.Context) error {

	res, err := handle.userBusiness.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get all data user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromCoreList(res),
	})

}

//	Menampilkan data user berdasarkan ID
func (handle *UserHandler) GetUserID(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "ID not recognized",
		})
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Unathorized",
		})
	}

	userId, err := handle.userBusiness.GetDatabyID(uint(idUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed get your data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    userId,
	})
}
