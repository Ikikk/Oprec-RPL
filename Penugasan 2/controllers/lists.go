package controllers

import (
	"Penugasan-2/config"
	"Penugasan-2/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InsertListsToDB(db *sql.DB, list models.Lists) (err error) {

	rows, err := db.Query("INSERT INTO lists(list_title,tags_id,description,checklist) VALUES($1,$2,$3,$4) RETURNING list_id,list_title,tags_id,description,checklist", list.Title, list.Tag, list.Description, list.Check)
	if err != nil {
		return err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := models.Lists{}

	// insert rows to result
	rows.Scan(&result.ID, &result.Title)
	return nil
}

func PostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ToDos = models.Lists{}
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&ToDos); err != nil {
				log.Fatal(err)
			}
		}

		err := InsertListsToDB(config.DB, ToDos)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

func GetListDB(db *sql.DB) ([]models.Lists, error) {
	var result []models.Lists

	rows, err := db.Query("select * from lists")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var list models.Lists
		err = rows.Scan(&list.ID, &list.Title, &list.Tag, &list.Description, &list.Check)
		if err != nil {
			panic(err)
		}
		fmt.Println(list.ID, list.Title, list.Tag, list.Description, list.Check)
		result = append(result, list)
	}
	return result, nil
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		list, err := GetListDB(config.DB)
		if err != nil {
			fmt.Println(err)
		}

		res, err := json.Marshal(list)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}

func DeleteListDB(db *sql.DB, id int) (err error) {
	errs := db.QueryRow("DELETE FROM lists WHERE list_id=$1", id)
	return errs.Err()
}

func DeleteList(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("%d", id)
		err = DeleteListDB(config.DB, id)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success deleted"))
	}
}
