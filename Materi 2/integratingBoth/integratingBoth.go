package integratingBoth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx"
)

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"user"`
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://postgres:root@localhost:5432/pgx")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertToDB(db *sql.DB, user User) (*User, error) {
	rows, err := db.Query("INSERT INTO users (name) VALUES ($1) RETURNING id, name", user.Name)
	if err != nil {
		return nil, err
	}

	// call rows.Next() to move pointer to first result set
	rows.Next()

	// result container object
	result := User{}

	// insert rows to result
	rows.Scan(&result.ID, &result.Name)
	return &result, nil
}

func GetAll(db *sql.DB) ([]User, error) {
	var result []User

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		result = append(result, user)
	}

	return result, nil
}

func main() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	res, err := GetAll(db)
	if err != nil {
		fmt.Println(err)
	}

	// turning result into JSON
	jsonMap := map[string]interface{}{
		"data": res,
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tell the client that the content type is json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(b))
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
