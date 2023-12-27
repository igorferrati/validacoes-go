package main

import (
	"github.com/igorferrati/api-go-gin/database"
	"github.com/igorferrati/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
