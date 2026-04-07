package repository

import (
	"context"
	"errors"
	"server/internal/domain/entities"
	"server/internal/interfaces/repository"
	"time"

	"gorm.io/gorm"
)

// ReservationModel représente le modèle GORM pour la persistance
type ReservationModel struct {
	ID           string    `gorm:"column:id;primaryKey;type:varchar(64)"`
	Nom          string    `gorm:"column:nom;type:varchar(255)"`
	Prenom       string    `gorm:"column:prenom;type:varchar(255)"`
	DocumentID   string    `gorm:"column:document_id;type:varchar(128)"`
	DocumentType string    `gorm:"column:document_type;type:varchar(128)"`
	Telephone    string    `gorm:"column:telephone;type:varchar(64)"`
	Email        string    `gorm:"column:email;type:varchar(255)"`
	DepartVille  string    `gorm:"column:depart_ville;type:varchar(255)"`
	ArriveeVille string    `gorm:"column:arrivee_ville;type:varchar(255)"`
	VolNumero    string    `gorm:"column:vol_numero;type:varchar(64)"`
	DateDepart   string    `gorm:"column:date_depart;type:varchar(64)"`
	DateArrivee  string    `gorm:"column:date_arrivee;type:varchar(64)"`
	Motif        string    `gorm:"column:motif;type:varchar(255)"`
	Bagages      string    `gorm:"column:bagages;type:varchar(64)"`
	Classe       string    `gorm:"column:classe;type:varchar(64)"`
	Statut       string    `gorm:"column:statut;type:varchar(64)"`
	Special      string    `gorm:"column:special;type:varchar(255)"`
	Notes        string    `gorm:"column:notes;type:text"`
	Route        string    `gorm:"column:route;type:varchar(255)"`
	Depart       string    `gorm:"column:depart;type:varchar(255)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ReservationModel) TableName() string {
	return "reservations"
}

// gormReservationRepository implémente ReservationRepository avec GORM
type gormReservationRepository struct {
	db *gorm.DB
}

// NewGormReservationRepository crée une nouvelle instance du repository GORM
func NewGormReservationRepository(db *gorm.DB) repository.ReservationRepository {
	return &gormReservationRepository{db: db}
}

// Save sauvegarde une nouvelle réservation
func (r *gormReservationRepository) Save(ctx context.Context, reservation *entities.Reservation) error {
	model := toReservationModel(reservation)
	return r.db.WithContext(ctx).Create(model).Error
}

// FindByID trouve une réservation par son ID
func (r *gormReservationRepository) FindByID(ctx context.Context, id entities.ReservationID) (*entities.Reservation, error) {
	var model ReservationModel
	err := r.db.WithContext(ctx).Where("id = ?", string(id)).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("réservation non trouvée")
		}
		return nil, err
	}
	return toReservationEntity(&model), nil
}

// FindAll retourne toutes les réservations
func (r *gormReservationRepository) FindAll(ctx context.Context) ([]*entities.Reservation, error) {
	var models []ReservationModel
	err := r.db.WithContext(ctx).Order("created_at desc").Find(&models).Error
	if err != nil {
		return nil, err
	}

	var reservations []*entities.Reservation
	for _, model := range models {
		reservations = append(reservations, toReservationEntity(&model))
	}
	return reservations, nil
}

// Update met à jour une réservation existante
func (r *gormReservationRepository) Update(ctx context.Context, reservation *entities.Reservation) error {
	model := toReservationModel(reservation)
	return r.db.WithContext(ctx).Save(model).Error
}

// Delete supprime une réservation par son ID
func (r *gormReservationRepository) Delete(ctx context.Context, id entities.ReservationID) error {
	return r.db.WithContext(ctx).Delete(&ReservationModel{}, "id = ?", string(id)).Error
}

// toReservationModel convertit une entité domaine vers le modèle GORM
func toReservationModel(r *entities.Reservation) *ReservationModel {
	return &ReservationModel{
		ID:           string(r.ID),
		Nom:          r.Nom,
		Prenom:       r.Prenom,
		DocumentID:   r.DocumentID,
		DocumentType: r.DocumentType,
		Telephone:    r.Telephone,
		Email:        r.Email,
		DepartVille:  r.DepartVille,
		ArriveeVille: r.ArriveeVille,
		VolNumero:    r.VolNumero,
		DateDepart:   r.DateDepart,
		DateArrivee:  r.DateArrivee,
		Motif:        r.Motif,
		Bagages:      r.Bagages,
		Classe:       r.Classe,
		Statut:       r.Statut,
		Special:      r.Special,
		Notes:        r.Notes,
		Route:        r.Route,
		Depart:       r.Depart,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}

// toReservationEntity convertit un modèle GORM vers l'entité domaine
func toReservationEntity(m *ReservationModel) *entities.Reservation {
	return &entities.Reservation{
		ID:           entities.ReservationID(m.ID),
		Nom:          m.Nom,
		Prenom:       m.Prenom,
		DocumentID:   m.DocumentID,
		DocumentType: m.DocumentType,
		Telephone:    m.Telephone,
		Email:        m.Email,
		DepartVille:  m.DepartVille,
		ArriveeVille: m.ArriveeVille,
		VolNumero:    m.VolNumero,
		DateDepart:   m.DateDepart,
		DateArrivee:  m.DateArrivee,
		Motif:        m.Motif,
		Bagages:      m.Bagages,
		Classe:       m.Classe,
		Statut:       m.Statut,
		Special:      m.Special,
		Notes:        m.Notes,
		Route:        m.Route,
		Depart:       m.Depart,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}