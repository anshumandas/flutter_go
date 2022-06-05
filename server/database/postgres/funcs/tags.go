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

func FetchTags() *[]models.Tag {
	var tags []models.Tag
	tx := postgres.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID")
	}).Find(&tags)
	if tx.Error != nil {
		fmt.Printf("error occurred with fetch tags function %v", tx.Error)
	}
	return &tags
}

func FetchLastTag() models.Tag {
	postgres.DB.First(&tag)
	return tag
}

func AddTag(name, description string, tags []string) string {
	newTag := models.Tag{Name: name}
	postgres.DB.Create(&newTag)
	return newTag.Name
}

func DeleteTag(id int) {
	postgres.DB.Delete(&tag, id)
}

func UpdateTag(id int, title, description, author string) {
	postgres.DB.Model(&tag).Where("id=?", id).Updates(map[string]interface{}{
		"title":       title,
		"description": description,
		"author":      author,
	})
}
