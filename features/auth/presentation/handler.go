package presentation

import (
	"net/http"
	"construct-week1/features/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authHandler auth.Business
}

// Dependency Injection
func NewUserHandler(auth auth.Business) *AuthHandler {
	return &AuthHandler{
		authHandler: auth,
	}
}

func (handle *AuthHandler) Login(c echo.Context) error {
	authData := auth.User{}
	c.Bind(&authData)

	token, e := handle.authHandler.LoginUsers(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Email or Password incorrect",
		})
	}
	data := map[string]interface{}{
		"token": token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
		"data":    data,
	})
}
