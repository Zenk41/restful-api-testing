package models

import "github.com/go-playground/validator/v10"

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (input *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}