package migrations

import (
	"backend/config"
	"backend/models"
	"log"
)

//Migrate DB structures
func Migrate() {
	log.Println("Initiating migration...")
	err := config.GetDbConnection().AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Category{},
		&models.Shop{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
