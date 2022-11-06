package tokens

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var storageTokenAdmin []string = make([]string, 5)

func CreateTokenAdmin(adminID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": adminID,
		"expire":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokeString, err := token.SignedString([]byte("secretkey"))

	if err != nil {
		log.Fatalf("error when creating token: %v", err)
	}
	storageTokenAdmin = append(storageTokenAdmin, tokeString)
	return tokeString
}

func CheckTokenAdmin(token string) bool {
	for _, tkn := range storageTokenAdmin {
		if tkn == token {
			return true
		}
	}
	return false
}

func ExtractTokenAdmin(e echo.Context) uint {
	Admin := e.Get("user").(*jwt.Token)
	isListed := CheckTokenAdmin(Admin.Raw)

	if !isListed {
		return 0
	}

	if Admin.Valid {
		claims := Admin.Claims.(jwt.MapClaims)
		adminID := claims["adminID"].(float64)
		return uint(adminID)
	}

	return 0
}

func LogoutAdmin(token string) bool {
	for idx, tkn := range storageTokenAdmin {
		if tkn == token {
			storageTokenAdmin = append(storageTokenAdmin[:idx], storageTokenAdmin[idx+1:]...)
		}
	}
	return true
}
