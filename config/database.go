package config

import (
	"fmt"
	"log"
	"reffil_finder/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/reffil_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Promotion{},
		&models.Cart{},
		&models.Transaksi{},
		&models.DetailTransaksi{},
	)
	if err != nil {
		log.Fatalf("Error in migration: %v", err)
	}

	fmt.Println("Database connected and migrated")
}
