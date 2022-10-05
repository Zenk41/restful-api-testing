package models

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (input *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}