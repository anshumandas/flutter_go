package models

type Tag struct {
	Name string `gorm:"primaryKey" form:"name" json:"name" binding:"required"`
}
