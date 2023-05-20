package main

import (
	"fmt"

	"github.com/nishiduka/rest-api/initializers"
	"github.com/nishiduka/rest-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	fmt.Println("Initilizing Migrations")
	initializers.DB.AutoMigrate(&models.Post{})

	fmt.Println("Finish")
}
