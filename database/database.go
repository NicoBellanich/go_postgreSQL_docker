package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nicobellanich/go_postgreSQL_docker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBase struct {
	Instance *gorm.DB
}

var DB DataBase

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	connectedDatabase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	connectedDatabase.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	connectedDatabase.AutoMigrate(&models.Fact{})

	DB = DataBase{
		Instance: connectedDatabase,
	}
}
