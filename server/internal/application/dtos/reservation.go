package dtos

// CreateReservationRequest représente la requête de création de réservation
type CreateReservationRequest struct {
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

// UpdateReservationRequest représente la requête de mise à jour de réservation
type UpdateReservationRequest struct {
	Nom          string `json:"nom"`
	Prenom       string `json:"prenom"`
	DocumentID   string `json:"document_id"`
	DocumentType string `json:"document_type"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	DepartVille  string `json:"depart_ville"`
	ArriveeVille string `json:"arrivee_ville"`
	VolNumero    string `json:"vol_numero"`
	DateDepart   string `json:"date_depart"`
	DateArrivee  string `json:"date_arrivee"`
	Motif        string `json:"motif"`
	Bagages      string `json:"bagages"`
	Classe       string `json:"classe"`
	Statut       string `json:"statut"`
	Special      string `json:"special"`
	Notes        string `json:"notes"`
	Route        string `json:"route"`
	Depart       string `json:"depart"`
}

// ReservationResponse représente la réponse contenant une réservation
type ReservationResponse struct {
	ID           string `json:"id"`
	Nom          string `json:"nom"`
	Prenom       string `json:"prenom"`
	DocumentID   string `json:"document_id"`
	DocumentType string `json:"document_type"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	DepartVille  string `json:"depart_ville"`
	ArriveeVille string `json:"arrivee_ville"`
	VolNumero    string `json:"vol_numero"`
	DateDepart   string `json:"date_depart"`
	DateArrivee  string `json:"date_arrivee"`
	Motif        string `json:"motif"`
	Bagages      string `json:"bagages"`
	Classe       string `json:"classe"`
	Statut       string `json:"statut"`
	Special      string `json:"special"`
	Notes        string `json:"notes"`
	Route        string `json:"route"`
	Depart       string `json:"depart"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// ReservationsResponse représente la réponse contenant une liste de réservations
type ReservationsResponse struct {
	Reservations []ReservationResponse `json:"data"`
	Status       int                   `json:"status"`
	Message      string                `json:"message"`
}