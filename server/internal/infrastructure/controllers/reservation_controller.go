package controllers

import (
	"net/http"
	"server/internal/application/dtos"
	"server/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

// ReservationController gère les requêtes HTTP pour les réservations
type ReservationController struct {
	createReservationUseCase *usecases.CreateReservationUseCase
	listReservationsUseCase  *usecases.ListReservationsUseCase
	updateReservationUseCase *usecases.UpdateReservationUseCase
	deleteReservationUseCase *usecases.DeleteReservationUseCase
}

// NewReservationController crée une nouvelle instance du contrôleur
func NewReservationController(
	createUC *usecases.CreateReservationUseCase,
	listUC *usecases.ListReservationsUseCase,
	updateUC *usecases.UpdateReservationUseCase,
	deleteUC *usecases.DeleteReservationUseCase,
) *ReservationController {
	return &ReservationController{
		createReservationUseCase: createUC,
		listReservationsUseCase:  listUC,
		updateReservationUseCase: updateUC,
		deleteReservationUseCase: deleteUC,
	}
}

// ListReservations gère la requête GET /reservations
func (ctrl *ReservationController) ListReservations(c *gin.Context) {
	response, err := ctrl.listReservationsUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to load reservations",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateReservation gère la requête POST /reservations
func (ctrl *ReservationController) CreateReservation(c *gin.Context) {
	var request dtos.CreateReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid reservation data",
		})
		return
	}

	response, err := ctrl.createReservationUseCase.Execute(c.Request.Context(), request)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "nom et prénom sont requis" ||
		   err.Error() == "ville de départ et d'arrivée sont requises" ||
		   err.Error() == "motif est requis" ||
		   err.Error() == "classe est requise" ||
		   err.Error() == "statut est requis" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{
			"status":  status,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Reservation created",
		"data":    response,
	})
}

// UpdateReservation gère la requête PUT /reservations/:id
func (ctrl *ReservationController) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var request dtos.UpdateReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid reservation data",
		})
		return
	}

	response, err := ctrl.updateReservationUseCase.Execute(c.Request.Context(), id, request)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "réservation non trouvée" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"status":  status,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Reservation updated",
		"data":    response,
	})
}

// DeleteReservation gère la requête DELETE /reservations/:id
func (ctrl *ReservationController) DeleteReservation(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.deleteReservationUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "réservation non trouvée" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"status":  status,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Reservation deleted",
	})
}