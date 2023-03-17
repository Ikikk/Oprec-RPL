package config

import (
	"fmt"
	"log"
	"os"

	"Penugasan-3/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	pgdb := os.Getenv("PGDATABASE")

	data := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, pgdb,
	)

	db, err := gorm.Open(postgres.Open(data), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.Debug().AutoMigrate(models.Customers{}, models.Accounts{})
	return db, nil
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	dbSQL.Close()

	return
}
