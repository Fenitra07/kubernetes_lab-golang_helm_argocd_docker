package usecases

import (
	"context"
	"errors"
	"server/internal/application/dtos"
	"server/internal/domain/entities"
	"server/internal/interfaces/repository"
)

// UpdateReservationUseCase gère la mise à jour de réservations
type UpdateReservationUseCase struct {
	reservationRepo repository.ReservationRepository
}

// NewUpdateReservationUseCase crée une nouvelle instance du use case
func NewUpdateReservationUseCase(repo repository.ReservationRepository) *UpdateReservationUseCase {
	return &UpdateReservationUseCase{
		reservationRepo: repo,
	}
}

// Execute exécute le use case de mise à jour de réservation
func (uc *UpdateReservationUseCase) Execute(ctx context.Context, id string, req dtos.UpdateReservationRequest) (*dtos.ReservationResponse, error) {
	reservationID := entities.ReservationID(id)

	// Vérification de l'existence de la réservation
	existingReservation, err := uc.reservationRepo.FindByID(ctx, reservationID)
	if err != nil {
		return nil, errors.New("réservation non trouvée")
	}

	// Mise à jour des informations
	if req.Nom != "" && req.Prenom != "" {
		existingReservation.UpdateMetadonnées(req.Nom, req.Prenom, req.Telephone, req.Email)
	}

	if req.DocumentID != "" {
		existingReservation.UpdateDocuments(req.DocumentID, req.DocumentType)
	}

	if req.VolNumero != "" {
		existingReservation.UpdateVol(req.VolNumero, req.DateDepart, req.DateArrivee)
	}

	existingReservation.UpdateDetails(
		req.Motif,
		req.Bagages,
		req.Classe,
		req.Statut,
		req.Special,
		req.Notes,
		req.Route,
		req.Depart,
	)

	// Sauvegarde via le repository
	if err := uc.reservationRepo.Update(ctx, existingReservation); err != nil {
		return nil, err
	}

	// Conversion vers DTO de réponse
	response := &dtos.ReservationResponse{
		ID:           existingReservation.ID.String(),
		Nom:          existingReservation.Nom,
		Prenom:       existingReservation.Prenom,
		DocumentID:   existingReservation.DocumentID,
		DocumentType: existingReservation.DocumentType,
		Telephone:    existingReservation.Telephone,
		Email:        existingReservation.Email,
		DepartVille:  existingReservation.DepartVille,
		ArriveeVille: existingReservation.ArriveeVille,
		VolNumero:    existingReservation.VolNumero,
		DateDepart:   existingReservation.DateDepart,
		DateArrivee:  existingReservation.DateArrivee,
		Motif:        existingReservation.Motif,
		Bagages:      existingReservation.Bagages,
		Classe:       existingReservation.Classe,
		Statut:       existingReservation.Statut,
		Special:      existingReservation.Special,
		Notes:        existingReservation.Notes,
		Route:        existingReservation.Route,
		Depart:       existingReservation.Depart,
		CreatedAt:    existingReservation.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    existingReservation.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return response, nil
}