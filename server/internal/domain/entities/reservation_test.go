package entities

import (
	"testing"
	"time"
)

func TestNewReservationInitializesFields(t *testing.T) {
	reservation := NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé")

	if reservation.ID == "" {
		t.Fatal("expected reservation ID to be generated")
	}

	if reservation.Nom != "Jean" {
		t.Errorf("expected Nom 'Jean', got %q", reservation.Nom)
	}

	if reservation.DepartVille != "TNR" {
		t.Errorf("expected DepartVille 'TNR', got %q", reservation.DepartVille)
	}

	if reservation.CreatedAt.IsZero() {
		t.Fatal("expected CreatedAt to be set")
	}
}

func TestUpdateDetailsUpdatesFields(t *testing.T) {
	reservation := NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé")
	initialUpdatedAt := reservation.UpdatedAt

	time.Sleep(2 * time.Millisecond)
	reservation.UpdateDetails("Affaires", "2", "Premium", "Confirmé", "Repas spécial", "Note de test", "TNR → NBO", "Terminal 2")

	if reservation.Motif != "Affaires" {
		t.Errorf("expected Motif 'Affaires', got %q", reservation.Motif)
	}

	if reservation.Classe != "Premium" {
		t.Errorf("expected Classe 'Premium', got %q", reservation.Classe)
	}

	if reservation.Route != "TNR → NBO" {
		t.Errorf("expected Route 'TNR → NBO', got %q", reservation.Route)
	}

	if !reservation.UpdatedAt.After(initialUpdatedAt) {
		t.Fatalf("expected UpdatedAt to be updated, got %v, initial %v", reservation.UpdatedAt, initialUpdatedAt)
	}
}
