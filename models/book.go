package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title            string `json:"title" form:"title"`
	Author           string `json:"author" form:"author"`
	Publisher        string `json:"publisher" form:"publisher"`
	Publication_Year string `json:"publication_year" form:"publication_year"`
	ISBN             string `json:"isbn" form:"isbn"`
	NumberOfPage     string `json:"number_of_page" form:"number_of_page"`
	Language         string `json:"language" form:"language"`
}