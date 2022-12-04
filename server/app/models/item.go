package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Status      string   `json:"status"` //TODO: change to enum
	CreatedBy   string   `json:"createdBy"`
	Type        string   `json:"type"` //this need not be saved and instead can be taken from table key
}
