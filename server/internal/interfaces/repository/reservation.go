package repository

import (
	"context"
	"server/internal/domain/entities"
)

// ReservationRepository définit le contrat pour l'accès aux données de réservation
type ReservationRepository interface {
	// Save sauvegarde une nouvelle réservation
	Save(ctx context.Context, reservation *entities.Reservation) error

	// FindByID trouve une réservation par son ID
	FindByID(ctx context.Context, id entities.ReservationID) (*entities.Reservation, error)

	// FindAll retourne toutes les réservations
	FindAll(ctx context.Context) ([]*entities.Reservation, error)

	// Update met à jour une réservation existante
	Update(ctx context.Context, reservation *entities.Reservation) error

	// Delete supprime une réservation par son ID
	Delete(ctx context.Context, id entities.ReservationID) error
}