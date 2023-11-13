package dto

import (
	"backend/helpers"
	"encoding/json"

	validator "github.com/go-playground/validator/v10"
)

type RegisterForm struct{}

type RegisterDTO struct {
	Name     string `form:"name" json:"name" binding:"required,min=8,max=20"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=50"`
}

func (f RegisterForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your name"
		}
		return errMsg[0]
	case "min", "max":
		return "Your name should be between 8 to 20 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f RegisterForm) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your email"
		}
		return errMsg[0]
	case "min", "max", "email":
		return "Please enter a valid email"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f RegisterForm) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Please enter your password"
	case "min", "max":
		return "Your password should be between 8 and 50 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Register ...
func (f RegisterForm) Register(err error) helpers.Response {
	var response = new(helpers.Response).SetStatus(400)

	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return *response.SetMessage("Something went wrong, please try again later")
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return *response.SetMessage(f.Name(err.Tag()))
			}

			if err.Field() == "Email" {
				return *response.SetMessage(f.Email(err.Tag()))
			}

			if err.Field() == "Password" {
				return *response.SetMessage(f.Password(err.Tag()))
			}

		}
	default:
		return *response.SetMessage("Invalid request")
	}

	return *response.SetMessage("Something went wrong, please try again later")
}
