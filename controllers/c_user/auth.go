package cuser

import (
	"mini_project/app/tokens"
	muser "mini_project/models/m_user"
	"mini_project/service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var authUserController service.AuthUserService = service.NewAuthUser()

func RegisterUser(c echo.Context) error {
	userInput := new(muser.Register)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "invalid request",
		})
	}
	err := userInput.ValidatorRegister()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	user := authUserController.RegisterUser(*userInput)

	return c.JSON(http.StatusAccepted, user)
}

func LoginUser(c echo.Context) error {
	userInput := new(muser.Login)
	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}
	err := userInput.ValidatorLogin()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	token := authUserController.LoginUser(*userInput)
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func LogoutUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	isListed := tokens.CheckTokenUser(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "response.FailToken",
		})
	}
	tokens.LogoutUser(user.Raw)
	return c.JSON(http.StatusOK, map[string]string{
		"messege": "success logout",
	})

}
