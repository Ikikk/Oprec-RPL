package controllers

import (
	"Penugasan-2/models"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ListControllers struct {
	list *models.Lists
}

func InsertListsToDB(db *sql.DB, list models.Lists) (*ListControllers, error) {
	rows, err := db.Query("INSERT INTO lists (list_title, tags_id, description, checklist) VALUES ($1, $2, $3, $4) RETURNING list_title, tags_id, description, checklist", list.Title, list.Tag, list.Description, list.Check)
	if err != nil {
		return nil, err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := ListControllers{}

	// insert rows to result
	rows.Scan(&result.list.ID, &result.list.Title)
	return &result, nil
}
