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

func FetchUser(name string) models.User {
	postgres.DB.First(&user, name)
	return user
}

func AddUser(uname, fname, lname, email, phash, phone string) uint {
	newUser := models.User{UserName: uname, FirstName: fname, LastName: lname, Email: email}
	//TODO add status enum here. We can have a workflow initiated here
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
