package config

import (
	"database/sql"

	"Penugasan-2/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://root:5555@localhost:5432/pgx")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertToDB(db *sql.DB, user models.Lists) (*models.Lists, error) {
	rows, err := db.Query("INSERT INTO users (name) VALUES ($1) RETURNING id, name", user.Title)
	if err != nil {
		return nil, err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := models.Lists{}

	// insert rows to result
	rows.Scan(&result.ID, &result.Title)
	return &result, nil
}

func GetAll(db *sql.DB) ([]models.Lists, error) {
	var result []models.Lists

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.Lists
		rows.Scan(&user.ID, &user.Title)
		result = append(result, user)
	}

	return result, nil
}
