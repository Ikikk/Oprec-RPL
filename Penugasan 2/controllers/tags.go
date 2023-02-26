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

func InsertTagsToDB(db *sql.DB, tag models.Tags) (err error) {
	rows, err := db.Query("INSERT INTO tags (label) VALUES ($1) RETURNING tags_id,label", tag.Label)
	if err != nil {
		return err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := models.Tags{}

	// insert rows to result
	rows.Scan(&result.ID, &result.Label)

	return nil
}

func PostTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ToDos = models.Tags{}
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&ToDos); err != nil {
				log.Fatal(err)
			}
		}

		err := InsertTagsToDB(config.DB, ToDos)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

func GetTagDB(db *sql.DB) ([]models.Tags, error) {
	var result []models.Tags

	rows, err := db.Query("select * from tags")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag models.Tags
		err = rows.Scan(&tag.ID, &tag.Label)
		if err != nil {
			panic(err)
		}
		fmt.Println(tag.ID, tag.Label)
		result = append(result, tag)
	}
	fmt.Println(rows)
	return result, nil
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tag, err := GetTagDB(config.DB)
		if err != nil {
			fmt.Println(err)
		}

		res, err := json.Marshal(tag)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}

func DeleteTagDB(db *sql.DB, id int) (err error) {
	errs := db.QueryRow("DELETE FROM tags WHERE tags_id=$1", id)
	return errs.Err()
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Println(err)
		}

		err = DeleteTagDB(config.DB, id)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success deleted"))
	}
}
