package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique_index"`
	FullName     string
	Email        string `gorm:"type:varchar(100);unique_index"`
	Bio          string `gorm:"size:1024"`
	Image        string
	ImageThumb   string
	PasswordHash string `gorm:"not null"`
	// TODO: not sure if I need this here
	Subscriptions []Subscription
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
