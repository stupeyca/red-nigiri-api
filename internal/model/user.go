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

// Add `Create` to `User` struct.
func (user *User) Create() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Add `BeforeCreate` to `User` struct.
func (user *User) BeforeCreate(*gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	user.Email = string(user.Email)
	return nil
}

// Add `isPasswordValid` to `User` struct. Compare the raw password sent by the user and the hashed one saved in DB.
func (user *User) IsPasswordValid(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.Where("Email = ?", email).First(&user).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}
