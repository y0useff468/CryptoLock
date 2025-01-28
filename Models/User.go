package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}
