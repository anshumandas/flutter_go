package models

//This does not use gorm model as we do not require those fields and primaryKey is Name
type API struct {
	//TODO add ID which is a short form of the name later for storage optimization
	Name        string   `gorm:"primaryKey" redis:"name" form:"name" json:"name" binding:"required"`
	Description string   `redis:"desc" form:"description" json:"description"`
	Children    []string `redis:"children" form:"children" json:"children"`
}
