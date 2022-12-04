package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type UserRef struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
}
