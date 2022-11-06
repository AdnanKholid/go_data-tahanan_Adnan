package muser

import "github.com/go-playground/validator/v10"

type Register struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input *Register) ValidatorRegister() error {
	validator := validator.New()

	err := validator.Struct(input)

	return err
}

func (input *Login) ValidatorLogin() error {
	validator := validator.New()

	err := validator.Struct(input)

	return err
}
