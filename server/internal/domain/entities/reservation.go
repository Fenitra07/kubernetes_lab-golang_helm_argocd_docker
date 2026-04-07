package entities

import (
	"fmt"
	"time"
)

// ReservationID représente un identifiant unique de réservation
type ReservationID string

// NewReservationID génère un nouvel ID de réservation
func NewReservationID() ReservationID {
	return ReservationID(fmt.Sprintf("RES-%d", time.Now().UnixNano()))
}

// String retourne la représentation string de l'ID
func (id ReservationID) String() string {
	return string(id)
}

// Reservation représente une réservation de vol dans le domaine métier
type Reservation struct {
	ID           ReservationID
	Nom          string
	Prenom       string
	DocumentID   string
	DocumentType string
	Telephone    string
	Email        string
	DepartVille  string
	ArriveeVille string
	VolNumero    string
	DateDepart   string
	DateArrivee  string
	Motif        string
	Bagages      string
	Classe       string
	Statut       string
	Special      string
	Notes        string
	Route        string
	Depart       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewReservation crée une nouvelle réservation
func NewReservation(nom, prenom, departVille, arriveeVille, motif, classe, statut string) *Reservation {
	now := time.Now()
	return &Reservation{
		ID:           NewReservationID(),
		Nom:          nom,
		Prenom:       prenom,
		DepartVille:  departVille,
		ArriveeVille: arriveeVille,
		Motif:        motif,
		Classe:       classe,
		Statut:       statut,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// UpdateMetadonnées met à jour les informations personnelles
func (r *Reservation) UpdateMetadonnées(nom, prenom, telephone, email string) {
	r.Nom = nom
	r.Prenom = prenom
	r.Telephone = telephone
	r.Email = email
	r.UpdatedAt = time.Now()
}

// UpdateDocuments met à jour les informations de document
func (r *Reservation) UpdateDocuments(documentID, documentType string) {
	r.DocumentID = documentID
	r.DocumentType = documentType
	r.UpdatedAt = time.Now()
}

// UpdateVol met à jour les informations de vol
func (r *Reservation) UpdateVol(volNumero, dateDepart, dateArrivee string) {
	r.VolNumero = volNumero
	r.DateDepart = dateDepart
	r.DateArrivee = dateArrivee
	r.UpdatedAt = time.Now()
}

// UpdateDetails met à jour les détails de la réservation
func (r *Reservation) UpdateDetails(motif, bagages, classe, statut, special, notes, route, depart string) {
	r.Motif = motif
	r.Bagages = bagages
	r.Classe = classe
	r.Statut = statut
	r.Special = special
	r.Notes = notes
	r.Route = route
	r.Depart = depart
	r.UpdatedAt = time.Now()
}