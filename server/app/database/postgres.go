package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/flutter_go/app/models"
	pg "github.com/flutter_go/framework/base/databases/postgres"
	"github.com/flutter_go/settings"
	"gorm.io/driver/postgres" //postgres var is from here
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgre() {
	var err error

	var newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	var s = settings.PostgresSettings

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		s.Host, s.User, s.Password, s.DBName, s.Addr, s.SSLMode, s.TimeZone,
	)

	pg.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
	})
	pg.CheckError(err)

	_ = pg.DB.AutoMigrate(
		&models.Item{},
	)

}
