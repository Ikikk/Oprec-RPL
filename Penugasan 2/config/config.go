package config

import (
	"database/sql"

	"Penugasan-2/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	DB *sql.DB
)

type ToDoConfig struct {
	toDo *models.Lists
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://postgres:5555@localhost:5432/toDo")
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
