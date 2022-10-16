package platform

import (
	"customers/src/pkg/domain/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDB(dbURL string) *gorm.DB {
	db := openDatabase(dbURL)
	migrate(db)
	return db
}

func openDatabase(dbURL string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("unable connect to db")
	}
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatal("unable migrate customer")
	}
}
