package service

import (
	muser "mini_project/models/m_user"
	"mini_project/repository"
	ruser "mini_project/repository/r_user"
)

type AuthUserService struct {
	repo repository.AuthRepo
}

func NewAuthUser() AuthUserService {
	return AuthUserService{
		repo: &ruser.AuthRepoImpl{},
	}
}

func (a *AuthUserService) RegisterUser(input muser.Register) muser.User {
	return a.repo.Register(input)
}

func (a *AuthUserService) LoginUser(input muser.Login) string {
	return a.repo.Login(input)
}
