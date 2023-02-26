package cmd

import (
	"Penugasan-2/config"
	"fmt"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	// alex := User{
	// 	Title: "alex",
	// }

	// val, err := config.InsertToDB(db, alex)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	res, err := config.GetAll(db)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println((val))
	fmt.Println(res)
}
