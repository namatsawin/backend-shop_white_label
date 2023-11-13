package services

import (
	userdto "backend/dto/user"
	"backend/helpers"
	"backend/models"
	"backend/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func (s UserService) GetUserByID(id string) helpers.Response {
	var response helpers.Response

	user, err := s.userRepository.GetUserByID(id)

	if user == nil {
		return *response.SetStatus(404).SetMessage("the user could not be found.").Build()
	}

	if err != nil {
		return *response.SetStatus(500).SetMessage(err.Error()).Build()
	}

	return *response.SetStatus(200).SetMessage("success").SetData(user)
}

func (s UserService) Register(form userdto.RegisterDTO) helpers.Response {
	var response helpers.Response

	err := s.userRepository.Create(&models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	})

	if err != nil {
		return *response.SetStatus(406).SetMessage(err.Error()).Build()
	}

	return *response.SetStatus(201).SetMessage("success").Build()
}
