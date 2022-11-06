package tokens

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var storageTokenStore []string = make([]string, 5)

func CreateTokenStore(storeID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"storeID": storeID,
		"expire":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokeString, err := token.SignedString([]byte("secretkey"))

	if err != nil {
		log.Fatalf("error when creating token: %v", err)
	}
	storageTokenStore = append(storageTokenStore, tokeString)
	return tokeString
}

func CheckTokenStore(token string) bool {
	for _, tkn := range storageTokenStore {
		if tkn == token {
			return true
		}
	}
	return false
}

func ExtractTokenStore(e echo.Context) uint {
	store := e.Get("user").(*jwt.Token)
	isListed := CheckTokenStore(store.Raw)

	if !isListed {
		return 0
	}

	if store.Valid {
		claims := store.Claims.(jwt.MapClaims)
		storeID := claims["storeID"].(float64)
		return uint(storeID)
	}

	return 0
}

func LogoutStore(token string) bool {
	for idx, tkn := range storageTokenStore {
		if tkn == token {
			storageTokenStore = append(storageTokenStore[:idx], storageTokenStore[idx+1:]...)
		}
	}
	return true
}
