package repositories

import (
	"backend/database"
	"backend/models"
	"errors"
)

type UserRepository struct{}

func (r *UserRepository) GetUserByID(id string) (user *models.User, err error) {
	db := database.GetDB().Model(&models.User{})

	result := db.Where("id = ?", id).First(&user)

	if user.Id == 0 {
		return nil, errors.New("the user could not be found")
	}

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	db := database.GetDB().Model(&models.User{})
	return db.Create(&user).Error
}
