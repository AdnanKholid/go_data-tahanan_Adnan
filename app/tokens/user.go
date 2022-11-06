package tokens

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var storageTokenUser []string = make([]string, 5)

func CreateTokenUser(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokeString, err := token.SignedString([]byte("secretkey"))

	if err != nil {
		log.Fatalf("error when creating token: %v", err)
	}
	storageTokenUser = append(storageTokenUser, tokeString)
	return tokeString
}

func CheckTokenUser(token string) bool {
	for _, tkn := range storageTokenUser {
		if tkn == token {
			return true
		}
	}
	return false
}

func ExtractTokenUser(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	isListed := CheckTokenUser(user.Raw)

	if !isListed {
		return 0
	}

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["userID"].(float64)
		return uint(userID)
	}

	return 0
}

func LogoutUser(token string) bool {
	for idx, tkn := range storageTokenUser {
		if tkn == token {
			storageTokenUser = append(storageTokenUser[:idx], storageTokenUser[idx+1:]...)
		}
	}
	return true
}
