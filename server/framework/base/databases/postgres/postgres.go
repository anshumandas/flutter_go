package postgres

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
