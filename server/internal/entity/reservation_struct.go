package entity

import "time"

type Reservation struct {
	ID           string    `json:"id" gorm:"column:id;primaryKey;type:varchar(64)"`
	Nom          string    `json:"nom" gorm:"column:nom;type:varchar(255)"`
	Prenom       string    `json:"prenom" gorm:"column:prenom;type:varchar(255)"`
	DocumentID   string    `json:"document_id" gorm:"column:document_id;type:varchar(128)"`
	DocumentType string    `json:"document_type" gorm:"column:document_type;type:varchar(128)"`
	Telephone    string    `json:"telephone" gorm:"column:telephone;type:varchar(64)"`
	Email        string    `json:"email" gorm:"column:email;type:varchar(255)"`
	DepartVille  string    `json:"depart_ville" gorm:"column:depart_ville;type:varchar(255)"`
	ArriveeVille string    `json:"arrivee_ville" gorm:"column:arrivee_ville;type:varchar(255)"`
	VolNumero    string    `json:"vol_numero" gorm:"column:vol_numero;type:varchar(64)"`
	DateDepart   string    `json:"date_depart" gorm:"column:date_depart;type:varchar(64)"`
	DateArrivee  string    `json:"date_arrivee" gorm:"column:date_arrivee;type:varchar(64)"`
	Motif        string    `json:"motif" gorm:"column:motif;type:varchar(255)"`
	Bagages      string    `json:"bagages" gorm:"column:bagages;type:varchar(64)"`
	Classe       string    `json:"classe" gorm:"column:classe;type:varchar(64)"`
	Statut       string    `json:"statut" gorm:"column:statut;type:varchar(64)"`
	Special      string    `json:"special" gorm:"column:special;type:varchar(255)"`
	Notes        string    `json:"notes" gorm:"column:notes;type:text"`
	Route        string    `json:"route" gorm:"column:route;type:varchar(255)"`
	Depart       string    `json:"depart" gorm:"column:depart;type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Reservation) TableName() string {
	return "reservations"
}

// ReservationRequest is used for create/update payloads.
type ReservationRequest struct {
	Nom          string `json:"nom" binding:"required"`
	Prenom       string `json:"prenom" binding:"required"`
	DocumentID   string `json:"document_id"`
	DocumentType string `json:"document_type"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	DepartVille  string `json:"depart_ville" binding:"required"`
	ArriveeVille string `json:"arrivee_ville" binding:"required"`
	VolNumero    string `json:"vol_numero"`
	DateDepart   string `json:"date_depart" binding:"required"`
	DateArrivee  string `json:"date_arrivee"`
	Motif        string `json:"motif" binding:"required"`
	Bagages      string `json:"bagages"`
	Classe       string `json:"classe" binding:"required"`
	Statut       string `json:"statut" binding:"required"`
	Special      string `json:"special"`
	Notes        string `json:"notes"`
	Route        string `json:"route"`
	Depart       string `json:"depart"`
}
