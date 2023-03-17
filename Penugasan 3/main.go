package main

import (
	"Penugasan-3/config"
	"Penugasan-3/controller"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	database, err := config.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer config.CloseDB(database)

	route := gin.Default()

	accountControl := controller.AccountControl{DB: database}
	customerControl := controller.CustomersControl{DB: database}

	route.POST("/account/post", accountControl.PostAccount)
	route.POST("/customer/post", customerControl.PostCustomer)
	route.PUT("/customer/:id/update", customerControl.UpdateCustomer)
	route.GET("/account/customer/:cust_id", accountControl.GetAllAccountByUser)
	route.DELETE("/account/customer/:cust_id/delete", accountControl.DeleteAccount)
	route.DELETE("/customer/:id/delete", customerControl.DeleteCustomer)

	fmt.Println("Running")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	route.Run(":" + port)
}
