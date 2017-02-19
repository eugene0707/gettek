package common

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(1)
	}

}

func CurrentDB() *gorm.DB {
	return db
}