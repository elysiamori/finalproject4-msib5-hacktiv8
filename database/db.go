package database

import (
	"fmt"
	"log"
	"os"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/admin"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func Config() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	config := ConfigDB{
		host:     os.Getenv("PGHOST"),
		port:     os.Getenv("PGPORT"),
		user:     os.Getenv("PGUSER"),
		password: os.Getenv("PGPASSWORD"),
		dbname:   os.Getenv("PGDATABASE"),
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.host, config.port, config.user, config.password, config.dbname)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	// migration
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.TransactionHistory{})

	admin.SeedAdmin(db)

	return db, nil
}
