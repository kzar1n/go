package main

import (
	"gin-api-rest/repository"
	"gin-api-rest/routes"
)

func main() {
	repository.OpenDB()
	routes.HandleRequests()
}
