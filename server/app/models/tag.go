package models

type Tag struct {
	Name string `gorm:"primaryKey" form:"name" json:"name" binding:"required"`
	//We can add top ten Tag usage etc. in case we need such access pattern. That way the search is reduced as this will get updated at EoD
}
