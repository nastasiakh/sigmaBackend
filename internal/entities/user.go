package entities

import "github.com/go-playground/validator/v10"

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName" validate:"required,min=3,max=20"`
	LastName  string `json:"lastName" validate:"required,min=3,max=20"`
	Email     string `json:"email" validate:"required,email"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}
