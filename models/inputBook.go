package models

import "github.com/go-playground/validator/v10"

type BookInput struct {
	Title            string `json:"title" form:"title" validate:"required"`
	Author           string `json:"author" form:"author" validate:"required"`
	Publisher        string `json:"publisher" form:"publisher" validate:"required"`
	Publication_Year string `json:"publication_year" form:"publication_year" validate:"required"`
	ISBN             string `json:"isbn" form:"isbn" validate:"required"`
	NumberOfPage     string `json:"number_of_page" form:"number_of_page" validate:"required"`
	Language         string `json:"language" form:"language" validate:"required"`
}

func (input *BookInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}