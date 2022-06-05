package postgrefuncs

import (
	"fmt"

	"github.com/flutter_go/database/postgres"
	"github.com/flutter_go/models"
	"gorm.io/gorm"
)

var (
	user models.User
)

func FetchUsers() *[]models.User {
	var users []models.User
	tx := postgres.DB.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID")
	}).Find(&users)
	if tx.Error != nil {
		fmt.Printf("error occurred with fetch users function %v", tx.Error)
	}
	return &users
}

func FetchLastUser() models.User {
	postgres.DB.First(&user)
	return user
}

func AddUser(name, description string, tags []string) uint {
	newUser := models.User{Name: name}
	postgres.DB.Create(&newUser)
	return newUser.ID
}

func DeleteUser(id int) {
	postgres.DB.Delete(&user, id)
}

func UpdateUser(id int, title, description, author string) {
	postgres.DB.Model(&user).Where("id=?", id).Updates(map[string]interface{}{
		"title":       title,
		"description": description,
		"author":      author,
	})
}
