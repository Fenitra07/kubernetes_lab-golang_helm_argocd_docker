package controler

import (
	"server/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LivraisonHandler interface {
	Login(*gin.Context)
	ListReservations(*gin.Context)
	CreateReservation(*gin.Context)
	UpdateReservation(*gin.Context)
	DeleteReservation(*gin.Context)
}

type livraisonHandler struct {
	db *gorm.DB
}

func NewHandler() LivraisonHandler {
	db := config.DatabaseConnex()

	return &livraisonHandler{db: db}
}
