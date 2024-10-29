package main

import (
	"fmt"
	"go-zip-code-temperature/config"
	"go-zip-code-temperature/internal/handler"
	"log"
	"net/http"
)

func main() {
	appConfig := getAppConfig()
	temperatureHandler := getHandler(appConfig)
	port := appConfig.WebServerPort
	http.HandleFunc("GET /temperature/{cep}", temperatureHandler.GetTemperature)

	fmt.Println("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getAppConfig() *config.Config {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading appConfig: %v", err)
	}
	return appConfig
}

func getHandler(config *config.Config) *handler.CityTemperatureHandler {
	return handler.NewCityTemperatureHandler(config)
}
