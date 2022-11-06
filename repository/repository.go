package repository

import (
	msel "mini_project/models/m_sel"
	msipir "mini_project/models/m_sipir"
	mtahanan "mini_project/models/m_tahanan"
	muser "mini_project/models/m_user"
)

type AuthRepo interface {
	Register(input muser.Register) muser.User
	Login(input muser.Login) string
}

type SelRepo interface {
	GetAll() []msel.Response
	GetByID(id string) msel.Response
	Create(input msel.Input) msel.Response
	Update(id string, input msel.Input) msel.Response
	Delete(id string) bool
}

type TahananRepo interface {
	GetAll() []mtahanan.Response
	GetByID(id string) mtahanan.Response
	Create(input mtahanan.Input) mtahanan.Response
	Update(id string, input mtahanan.Input) mtahanan.Response
	Delete(id string) bool
}

type SipirRepo interface {
	GetAll() []msipir.Response
	GetByID(id string) msipir.Response
	Create(input msipir.Input) msipir.Response
	Update(id string, input msipir.Input) msipir.Response
	Delete(id string) bool
}
