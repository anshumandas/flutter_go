package postgrefuncs

import (
	"fmt"

	"github.com/flutter_go/database/postgres"
	"github.com/flutter_go/models"
	"gorm.io/gorm"
)

var (
	tag models.Tag
)

func FetchTag(name string) models.Tag {
	postgres.DB.First(&tag, name)
	return tag
}

func FetchTags() *[]models.Tag {
	var tags []models.Tag
	tx := postgres.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("Name")
	}).Find(&tags)
	if tx.Error != nil {
		fmt.Printf("error occurred with fetch tags function %v", tx.Error)
	}
	return &tags
}

//These are called by the event consumers and not by REST services
func AddTag(name string) string {
	newTag := models.Tag{Name: name}
	postgres.DB.Create(&newTag)
	return newTag.Name
}

func AddTagLinks(name string, id []string) string {
	//TODO check if tag exists, if not then add else just append the ids
	newTag := models.Tag{Name: name}
	postgres.DB.Create(&newTag)
	return newTag.Name
}
