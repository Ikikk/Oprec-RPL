package controllers

import (
	"Penugasan-2/models"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TagControllers struct {
	tag *models.Tags
}

func ConnectTags() (*sql.DB, error) {
	db, err := sql.Open("toDo", "postgres://postgres:5555@localhost:5432/toDo")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertTagsToDB(db *sql.DB, tag *models.Tags) (*TagControllers, error) {
	rows, err := db.Query("INSERT INTO tags (label) VALUES ($1) RETURNING label", tag.Label)
	if err != nil {
		return nil, err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := TagControllers{}

	// insert rows to result
	rows.Scan(&result.tag.ID, &result.tag.Label)
	return &result, nil
}
