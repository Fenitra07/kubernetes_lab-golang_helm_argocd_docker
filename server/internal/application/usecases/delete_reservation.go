package usecases

import (
	"context"
	"errors"
	"server/internal/domain/entities"
	"server/internal/interfaces/repository"
)

// DeleteReservationUseCase gère la suppression de réservations
type DeleteReservationUseCase struct {
	reservationRepo repository.ReservationRepository
}

// NewDeleteReservationUseCase crée une nouvelle instance du use case
func NewDeleteReservationUseCase(repo repository.ReservationRepository) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{
		reservationRepo: repo,
	}
}

// Execute exécute le use case de suppression de réservation
func (uc *DeleteReservationUseCase) Execute(ctx context.Context, id string) error {
	reservationID := entities.ReservationID(id)

	// Vérification de l'existence de la réservation
	_, err := uc.reservationRepo.FindByID(ctx, reservationID)
	if err != nil {
		return errors.New("réservation non trouvée")
	}

	// Suppression via le repository
	return uc.reservationRepo.Delete(ctx, reservationID)
}