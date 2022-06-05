package postgrefuncs

import (
	"fmt"

	"github.com/flutter_go/database/postgres"
	"github.com/flutter_go/models"
	"gorm.io/gorm"
)

var (
	api models.API
)

func FetchAPIs() *[]models.API {
	var categories []models.API
	tx := postgres.DB.Preload("APIs", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID")
	}).Find(&categories)
	if tx.Error != nil {
		fmt.Printf("error occurred with fetch categories function %v", tx.Error)
	}
	return &categories
}

func FetchLastAPI() models.API {
	postgres.DB.First(&api)
	return api
}

func AddAPI(name, description string, children []string) string {
	newAPI := models.API{Name: name}
	postgres.DB.Create(&newAPI)
	return newAPI.Name
}

func DeleteAPI(id int) {
	postgres.DB.Delete(&api, id)
}

func UpdateAPI(id int, title, description, author string) {
	postgres.DB.Model(&api).Where("id=?", id).Updates(map[string]interface{}{
		"title":       title,
		"description": description,
		"author":      author,
	})
}
