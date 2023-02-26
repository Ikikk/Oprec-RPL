package config

import (
	"database/sql"

	"Penugasan-2/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ToDoConfig struct {
	toDo *models.Lists
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://root:5555@localhost:5432/pgx")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetAll(db *sql.DB) ([]ToDoConfig, error) {
	var list []ToDoConfig

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var toDoList ToDoConfig
		rows.Scan(&toDoList.toDo.Title, &toDoList.toDo.Tag, &toDoList.toDo.Description)
		list = append(list, toDoList)
	}

	return list, nil
}
