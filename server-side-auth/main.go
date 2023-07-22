package main

import (
	"example.com/models"
	"example.com/routes"
)

func main() {
	models.DbSetup()
	routes.RoutesSetup()
}
