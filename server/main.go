package main

import (
	"fmt"
	"log"
	"os"
	"server/internal/config"
	"server/internal/entity"
	"server/internal/route"
)

var ROOT_FOLDER string

func init() {
	if ROOT_FOLDER == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("unable to determine project root: %v", err)
		}
		ROOT_FOLDER = wd
	}
	os.Setenv(config.ROOT_FOLDER_VAR, ROOT_FOLDER)
	config.Load()
}

func main() {
	db := config.DatabaseConnex()
	if err := db.AutoMigrate(&entity.Reservation{}); err != nil {
		log.Fatalf("Failed to migrate reservation schema: %v", err)
	}

	apiAddress := fmt.Sprintf("%s:%d", config.AppConfiguration.API.Host, config.AppConfiguration.API.Port)
	if apiAddress == "" {
		log.Printf("EMPTY API_ADRESSSE: %s\n", apiAddress)
	}

	route.SetupRouter(apiAddress)
}
