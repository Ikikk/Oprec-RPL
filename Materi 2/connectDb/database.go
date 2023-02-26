package connectDb

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://root:5555@localhost:5432/pgx")
	if err != nil {
		return nil, err
	}
	return db, nil
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
}
