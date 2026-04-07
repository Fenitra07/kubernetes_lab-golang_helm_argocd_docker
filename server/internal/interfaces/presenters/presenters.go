package presenters

import "github.com/gin-gonic/gin"

// ReservationPresenter définit l'interface pour les contrôleurs de réservation
type ReservationPresenter interface {
	ListReservations(*gin.Context)
	CreateReservation(*gin.Context)
	UpdateReservation(*gin.Context)
	DeleteReservation(*gin.Context)
}

// AuthPresenter définit l'interface pour les contrôleurs d'authentification
type AuthPresenter interface {
	Login(*gin.Context)
}