package model

import (
	"red-nigiri-api/internal/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
The following values are automatically added by GORM:
- id
- CreatedAt
-	UpdatedAt
- DeletedAt
*/

type User struct {
	gorm.Model
	UUID      string `gorm:"type:uuid;default:gen_random_uuid()" json:"uuid"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	FirstName string `gorm:"size:255;not null" json:"firstName"`
	LastName  string `gorm:"size:255;not null" json:"lastName"`
	Password  string `gorm:"size:255;not null;" json:"-"`
}

func (user *User) Create() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) HashPassword(*gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	user.Email = string(user.Email)
	return nil
}
