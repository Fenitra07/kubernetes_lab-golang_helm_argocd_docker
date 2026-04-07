package usecases

import (
	"context"
	"server/internal/application/dtos"
	"server/internal/interfaces/repository"
)

// ListReservationsUseCase gère la récupération de toutes les réservations
type ListReservationsUseCase struct {
	reservationRepo repository.ReservationRepository
}

// NewListReservationsUseCase crée une nouvelle instance du use case
func NewListReservationsUseCase(repo repository.ReservationRepository) *ListReservationsUseCase {
	return &ListReservationsUseCase{
		reservationRepo: repo,
	}
}

// Execute exécute le use case de listage des réservations
func (uc *ListReservationsUseCase) Execute(ctx context.Context) (*dtos.ReservationsResponse, error) {
	// Récupération via le repository
	reservations, err := uc.reservationRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Conversion vers DTOs de réponse
	var reservationResponses []dtos.ReservationResponse
	for _, reservation := range reservations {
		response := dtos.ReservationResponse{
			ID:           reservation.ID.String(),
			Nom:          reservation.Nom,
			Prenom:       reservation.Prenom,
			DocumentID:   reservation.DocumentID,
			DocumentType: reservation.DocumentType,
			Telephone:    reservation.Telephone,
			Email:        reservation.Email,
			DepartVille:  reservation.DepartVille,
			ArriveeVille: reservation.ArriveeVille,
			VolNumero:    reservation.VolNumero,
			DateDepart:   reservation.DateDepart,
			DateArrivee:  reservation.DateArrivee,
			Motif:        reservation.Motif,
			Bagages:      reservation.Bagages,
			Classe:       reservation.Classe,
			Statut:       reservation.Statut,
			Special:      reservation.Special,
			Notes:        reservation.Notes,
			Route:        reservation.Route,
			Depart:       reservation.Depart,
			CreatedAt:    reservation.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:    reservation.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		reservationResponses = append(reservationResponses, response)
	}

	return &dtos.ReservationsResponse{
		Reservations: reservationResponses,
		Status:       200,
		Message:      "Reservations loaded",
	}, nil
}