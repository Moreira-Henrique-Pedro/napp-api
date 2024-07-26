package main

import (
	"log"

	"github.com/Moreira-Henrique-Pedro/napp-api/src/controller"
	"github.com/Moreira-Henrique-Pedro/napp-api/src/infra"
	"github.com/Moreira-Henrique-Pedro/napp-api/src/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	infra.Connect()
	defer infra.Disconnect()

	stockService := service.NewStockService()
	stockController := controller.NewStockController(stockService)

	stockController.InitRoutes()
}
