package usecases

import (
	"context"
	"errors"
	"server/internal/application/dtos"
	"server/internal/domain/entities"
	"server/internal/interfaces/repository"
)

// CreateReservationUseCase gère la création de réservations
type CreateReservationUseCase struct {
	reservationRepo repository.ReservationRepository
}

// NewCreateReservationUseCase crée une nouvelle instance du use case
func NewCreateReservationUseCase(repo repository.ReservationRepository) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		reservationRepo: repo,
	}
}

// Execute exécute le use case de création de réservation
func (uc *CreateReservationUseCase) Execute(ctx context.Context, req dtos.CreateReservationRequest) (*dtos.ReservationResponse, error) {
	// Validation métier
	if req.Nom == "" || req.Prenom == "" {
		return nil, errors.New("nom et prénom sont requis")
	}

	if req.DepartVille == "" || req.ArriveeVille == "" {
		return nil, errors.New("ville de départ et d'arrivée sont requises")
	}

	if req.Motif == "" {
		return nil, errors.New("motif est requis")
	}

	if req.Classe == "" {
		return nil, errors.New("classe est requise")
	}

	if req.Statut == "" {
		return nil, errors.New("statut est requis")
	}

	// Création de l'entité domaine
	reservation := entities.NewReservation(
		req.Nom,
		req.Prenom,
		req.DepartVille,
		req.ArriveeVille,
		req.Motif,
		req.Classe,
		req.Statut,
	)

	// Mise à jour des informations supplémentaires
	reservation.UpdateMetadonnées(req.Nom, req.Prenom, req.Telephone, req.Email)
	reservation.UpdateDocuments(req.DocumentID, req.DocumentType)
	reservation.UpdateVol(req.VolNumero, req.DateDepart, req.DateArrivee)
	reservation.UpdateDetails(req.Motif, req.Bagages, req.Classe, req.Statut, req.Special, req.Notes, req.Route, req.Depart)

	// Sauvegarde via le repository
	if err := uc.reservationRepo.Save(ctx, reservation); err != nil {
		return nil, err
	}

	// Conversion vers DTO de réponse
	response := &dtos.ReservationResponse{
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

	return response, nil
}