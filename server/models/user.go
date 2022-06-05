package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

type UserRef struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
}
