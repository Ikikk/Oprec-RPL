package main

import (
	"Penugasan-2/config"
	"Penugasan-2/controllers"
	"fmt"
	"net/http"
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

	http.HandleFunc("/lists/post", controllers.PostList)
	http.HandleFunc("/tags/post", controllers.PostTag)
	http.HandleFunc("/lists", controllers.GetList)
	http.HandleFunc("/tags", controllers.GetTag)
	http.HandleFunc("/lists/delete", controllers.DeleteList)
	http.HandleFunc("/tags/delete", controllers.DeleteTag)
	fmt.Println("Running")
	http.ListenAndServe(":8080", nil)
}
