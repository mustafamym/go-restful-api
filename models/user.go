package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err

	}
	user.Password = string(bytes)
	return nil
}

func (usr *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(providedPassword))
	if err != nil {
		return nil
	}
	return nil
}
