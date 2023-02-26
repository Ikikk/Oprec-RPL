package cmd

import (
	"Penugasan-2/config"
	"Penugasan-2/controllers"
	"Penugasan-2/models"
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

	toDoList := models.Lists{}
	toDoTag := models.Tags{}

	list, err := controllers.InsertListsToDB(db, toDoList)
	if err != nil {
		fmt.Println(err)
	}

	tag, err := controllers.InsertTagsToDB(db, toDoTag)
	if err != nil {
		fmt.Println(err)
	}

	res, err := config.GetAll(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(list)
	fmt.Println(tag)
	fmt.Println(res)
}
