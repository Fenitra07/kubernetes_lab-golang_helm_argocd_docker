package controler

import (
	"fmt"
	"net/http"
	"server/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *livraisonHandler) ListReservations(c *gin.Context) {
	var reservations []entity.Reservation
	if err := h.db.Order("created_at desc").Find(&reservations).Error; err != nil {
		fmt.Printf("ListReservations error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to load reservations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Reservations loaded",
		"data":    reservations,
	})
}

func (h *livraisonHandler) CreateReservation(c *gin.Context) {
	var request entity.ReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid reservation data",
		})
		return
	}

	newRes := entity.Reservation{
		ID:           fmt.Sprintf("RES-%d", time.Now().UnixNano()),
		Nom:          request.Nom,
		Prenom:       request.Prenom,
		DocumentID:   request.DocumentID,
		DocumentType: request.DocumentType,
		Telephone:    request.Telephone,
		Email:        request.Email,
		DepartVille:  request.DepartVille,
		ArriveeVille: request.ArriveeVille,
		VolNumero:    request.VolNumero,
		DateDepart:   request.DateDepart,
		DateArrivee:  request.DateArrivee,
		Motif:        request.Motif,
		Bagages:      request.Bagages,
		Classe:       request.Classe,
		Statut:       request.Statut,
		Special:      request.Special,
		Notes:        request.Notes,
		Route:        request.Route,
		Depart:       request.Depart,
	}

	if err := h.db.Create(&newRes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to save reservation",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Reservation created",
		"data":    newRes,
	})
}

func (h *livraisonHandler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var request entity.ReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid reservation data",
		})
		return
	}

	var existing entity.Reservation
	if err := h.db.Where("id = ?", id).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Reservation not found",
		})
		return
	}

	existing.Nom = request.Nom
	existing.Prenom = request.Prenom
	existing.DocumentID = request.DocumentID
	existing.DocumentType = request.DocumentType
	existing.Telephone = request.Telephone
	existing.Email = request.Email
	existing.DepartVille = request.DepartVille
	existing.ArriveeVille = request.ArriveeVille
	existing.VolNumero = request.VolNumero
	existing.DateDepart = request.DateDepart
	existing.DateArrivee = request.DateArrivee
	existing.Motif = request.Motif
	existing.Bagages = request.Bagages
	existing.Classe = request.Classe
	existing.Statut = request.Statut
	existing.Special = request.Special
	existing.Notes = request.Notes
	existing.Route = request.Route
	existing.Depart = request.Depart

	if err := h.db.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to update reservation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Reservation updated",
		"data":    existing,
	})
}

func (h *livraisonHandler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&entity.Reservation{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to delete reservation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Reservation deleted",
	})
}
