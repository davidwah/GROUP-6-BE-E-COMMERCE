package presentation

import (
	"net/http"

	"construct-week1/features/auth"
	"construct-week1/features/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authHandler auth.Business
}

// Dependency Injection
func NewAuthHandler(auth auth.Business) *AuthHandler {
	return &AuthHandler{
		authHandler: auth,
	}
}

func (handle *AuthHandler) Login(c echo.Context) error {
	authData := auth.User{}
	c.Bind(&authData)
	
	var input string = authData.Password
	token, e := handle.authHandler.LoginUsers(authData, input)
	if e != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Email or Password incorrect"))
	}

	data := map[string]interface{}{
		"token":    token,
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Login Success", data))
}
