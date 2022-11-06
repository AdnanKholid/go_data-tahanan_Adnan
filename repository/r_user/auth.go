package ruser

import (
	"mini_project/app/tokens"
	"mini_project/config"
	muser "mini_project/models/m_user"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepoImpl struct{}

func (a *AuthRepoImpl) Register(input muser.Register) muser.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser muser.User = muser.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
	}

	createdUser := muser.User{}

	result := config.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthRepoImpl) Login(input muser.Login) string {
	var user muser.User = muser.User{}

	config.DB.First(&user, "email=?", input.Email)

	if user.ID == 0 {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	token := tokens.CreateTokenUser(user.ID)

	return token
}
