package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// This struct will be used as like table in database.
type User struct {
	gorm.Model        // The gorm.Model specification adds some default properties to the Model, like id, created date, modified date, and deleted date.
	Name       string `json:"name"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
