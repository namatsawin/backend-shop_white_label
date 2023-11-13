package models

import (
	"backend/helpers"
	"errors"
	time "time"

	gorm "gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id, omitempty" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;size:50"`
	Email     string         `json:"email" gorm:"index:user_login_idx,unique;not null"`
	Password  string         `json:"password" gorm:"index:user_login_idx,not null"`
	Salt      string         `json:"salt, omitempty" gorm:"not null;size:20"`
	CreatedAt time.Time      `json:"created_at, omitempty" gorm:"default:NOW()"`
	UpdatedAt time.Time      `json:"updated_at, omitempty" gorm:"default:NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at, omitempty" gorm:"index:user_login_idx, where:deleted_at IS NOT NULL"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var encryptionHelper helpers.EncryptionHelper

	salt, err := encryptionHelper.GenerateRandomHex(20)

	if err != nil {
		return errors.New("cannot create user, please try again later")
	}

	hashedPassword, err := encryptionHelper.SHA256(u.Password + salt)

	if err != nil {
		return errors.New("cannot create user, please try again later")
	}

	u.Salt = salt
	u.Password = hashedPassword

	return nil
}
