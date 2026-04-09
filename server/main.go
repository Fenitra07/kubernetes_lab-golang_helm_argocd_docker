package main

import (
	"fmt"
	"log"
	"os"
	"server/internal/config"
	"server/internal/infrastructure/controllers"
	"server/internal/infrastructure/repository"
	"server/internal/infrastructure/routes"
	"server/internal/application/usecases"
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
	// Initialisation de la base de données
	db := config.DatabaseConnex()
	if err := db.AutoMigrate(&repository.ReservationModel{}, &repository.UserModel{}); err != nil {
		log.Fatalf("Failed to migrate reservation schema: %v", err)
	}

	// Initialisation des repositories (couche infrastructure)
	reservationRepo := repository.NewGormReservationRepository(db)
	userRepo := repository.NewGormUserRepository(db)

	// Initialisation des use cases (couche application)
	createReservationUC := usecases.NewCreateReservationUseCase(reservationRepo)
	listReservationsUC := usecases.NewListReservationsUseCase(reservationRepo)
	updateReservationUC := usecases.NewUpdateReservationUseCase(reservationRepo)
	deleteReservationUC := usecases.NewDeleteReservationUseCase(reservationRepo)
	loginUC := usecases.NewLoginUseCase(userRepo)

	// Initialisation des contrôleurs (couche infrastructure/interface adapters)
	reservationCtrl := controllers.NewReservationController(
		createReservationUC,
		listReservationsUC,
		updateReservationUC,
		deleteReservationUC,
	)
	authCtrl := controllers.NewAuthController(loginUC)

	// Initialisation du routeur
	router := routes.NewRouter(reservationCtrl, authCtrl)

	// Configuration de l'adresse API
	apiAddress := fmt.Sprintf("%s:%d", config.AppConfiguration.API.Host, config.AppConfiguration.API.Port)
	if apiAddress == "" {
		log.Printf("EMPTY API_ADRESSSE: %s\n", apiAddress)
	}

	// Démarrage du serveur
	router.SetupRouter(apiAddress)
}
