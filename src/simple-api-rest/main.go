package main

import (
	"fmt"
	"simple-rest-api/database"
	"simple-rest-api/routes"
)

func main() {

	database.OpenDB()

	fmt.Println("Iniciando...")
	routes.HandleRequest()
}
