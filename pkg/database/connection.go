package database

import (
	"log"
	"os"
	"time"

	"go-templ/infra/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
		},
	)

	dsn := os.Getenv("DB_URI")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Connected to database")

	db.AutoMigrate(&models.User{}, &models.Document{})
	log.Println("Database migrated")

	DB = db
}
