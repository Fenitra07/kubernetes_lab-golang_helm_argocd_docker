package usecases

import (
	"context"
	"errors"
	"server/internal/application/dtos"
	"server/internal/domain/entities"
	"testing"
)

type mockReservationRepo struct {
	saved       *entities.Reservation
	updated     *entities.Reservation
	deletedID   entities.ReservationID
	findAll     []*entities.Reservation
	findByID    *entities.Reservation
	saveErr     error
	findErr     error
	updateErr   error
	deleteErr   error
}

func (m *mockReservationRepo) Save(ctx context.Context, reservation *entities.Reservation) error {
	if m.saveErr != nil {
		return m.saveErr
	}
	m.saved = reservation
	return nil
}

func (m *mockReservationRepo) FindByID(ctx context.Context, id entities.ReservationID) (*entities.Reservation, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	return m.findByID, nil
}

func (m *mockReservationRepo) FindAll(ctx context.Context) ([]*entities.Reservation, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	return m.findAll, nil
}

func (m *mockReservationRepo) Update(ctx context.Context, reservation *entities.Reservation) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	m.updated = reservation
	return nil
}

func (m *mockReservationRepo) Delete(ctx context.Context, id entities.ReservationID) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	m.deletedID = id
	return nil
}

type mockUserRepo struct {
	user  *entities.User
	err   error
}

func (m *mockUserRepo) FindByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.user == nil {
		return nil, errors.New("utilisateur non trouvé")
	}
	if m.user.Email != email || m.user.Password != password {
		return nil, errors.New("identifiants invalides")
	}
	return m.user, nil
}

func (m *mockUserRepo) FindByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.user, nil
}

func TestCreateReservationUseCase_Success(t *testing.T) {
	repo := &mockReservationRepo{}
	uc := NewCreateReservationUseCase(repo)

	req := dtos.CreateReservationRequest{
		Nom:          "Jean",
		Prenom:       "Dupont",
		DepartVille:  "TNR",
		ArriveeVille: "NBO",
		Motif:        "Études",
		Classe:       "Economique",
		Statut:       "Confirmé",
		Telephone:    "0123456789",
		Email:        "jean.dupont@example.com",
		VolNumero:    "MK001",
		DateDepart:   "2026-05-01",
		DateArrivee:  "2026-05-02",
		Bagages:      "1",
		Special:      "Aucun",
		Notes:        "Préférence hublot",
		Route:        "TNR → NBO",
		Depart:       "Terminal 1",
	}

	response, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID == "" {
		t.Fatal("expected reservation ID to be set")
	}

	if repo.saved == nil {
		t.Fatal("expected reservation to be saved")
	}

	if repo.saved.Nom != req.Nom {
		t.Errorf("expected name %q, got %q", req.Nom, repo.saved.Nom)
	}
}

func TestCreateReservationUseCase_ValidationError(t *testing.T) {
	repo := &mockReservationRepo{}
	uc := NewCreateReservationUseCase(repo)

	req := dtos.CreateReservationRequest{
		Nom:          "",
		Prenom:       "Dupont",
		DepartVille:  "TNR",
		ArriveeVille: "NBO",
		Motif:        "Études",
		Classe:       "Economique",
		Statut:       "Confirmé",
	}

	_, err := uc.Execute(context.Background(), req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestCreateReservationUseCase_SaveError(t *testing.T) {
	repo := &mockReservationRepo{saveErr: errors.New("database failed")}
	uc := NewCreateReservationUseCase(repo)

	req := dtos.CreateReservationRequest{
		Nom:          "Jean",
		Prenom:       "Dupont",
		DepartVille:  "TNR",
		ArriveeVille: "NBO",
		Motif:        "Études",
		Classe:       "Economique",
		Statut:       "Confirmé",
	}

	_, err := uc.Execute(context.Background(), req)
	if err == nil {
		t.Fatal("expected save error, got nil")
	}
}

func TestListReservationsUseCase_ReturnsReservations(t *testing.T) {
	reservations := []*entities.Reservation{
		entities.NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé"),
	}
	repo := &mockReservationRepo{findAll: reservations}
	uc := NewListReservationsUseCase(repo)

	response, err := uc.Execute(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(response.Reservations) != 1 {
		t.Fatalf("expected 1 reservation, got %d", len(response.Reservations))
	}

	if response.Reservations[0].Nom != "Jean" {
		t.Errorf("expected reservation name Jean, got %q", response.Reservations[0].Nom)
	}
}

func TestLoginUseCase_Success(t *testing.T) {
	user := &entities.User{
		Email:    "test@example.com",
		Password: "secret",
	}
	repo := &mockUserRepo{user: user}
	uc := NewLoginUseCase(repo)

	request := dtos.LoginRequest{Login: "test@example.com", Password: "secret"}
	response, err := uc.Execute(context.Background(), request)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.Email != user.Email {
		t.Errorf("expected email %q, got %q", user.Email, response.Email)
	}
}

func TestLoginUseCase_InvalidCredentials(t *testing.T) {
	repo := &mockUserRepo{user: &entities.User{Email: "test@example.com", Password: "secret"}}
	uc := NewLoginUseCase(repo)

	request := dtos.LoginRequest{Login: "test@example.com", Password: "wrong"}
	_, err := uc.Execute(context.Background(), request)
	if err == nil {
		t.Fatal("expected invalid credentials error, got nil")
	}
}

func TestUpdateReservationUseCase_Success(t *testing.T) {
	reservation := entities.NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé")
	repo := &mockReservationRepo{findByID: reservation}
	uc := NewUpdateReservationUseCase(repo)

	req := dtos.UpdateReservationRequest{
		Motif:     "Affaires",
		Classe:    "Premium",
		Statut:    "Confirmé",
		Notes:     "Mise à jour",
		Route:     "TNR → NBO",
		Depart:    "Terminal 2",
		Telephone: "0123456789",
		Email:     "jean.dupont@example.com",
	}

	response, err := uc.Execute(context.Background(), reservation.ID.String(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.Motif != "Affaires" {
		t.Errorf("expected motif Affaires, got %q", response.Motif)
	}

	if repo.updated == nil {
		t.Fatal("expected reservation to be updated")
	}

	if repo.updated.Classe != "Premium" {
		t.Errorf("expected classe Premium, got %q", repo.updated.Classe)
	}
}

func TestUpdateReservationUseCase_NotFound(t *testing.T) {
	repo := &mockReservationRepo{findErr: errors.New("réservation non trouvée")}
	uc := NewUpdateReservationUseCase(repo)

	req := dtos.UpdateReservationRequest{Motif: "Affaires"}
	_, err := uc.Execute(context.Background(), "RES-123", req)
	if err == nil {
		t.Fatal("expected not found error, got nil")
	}
}

func TestDeleteReservationUseCase_Success(t *testing.T) {
	reservation := entities.NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé")
	repo := &mockReservationRepo{findByID: reservation}
	uc := NewDeleteReservationUseCase(repo)

	err := uc.Execute(context.Background(), reservation.ID.String())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if repo.deletedID != reservation.ID {
		t.Fatalf("expected deleted ID %q, got %q", reservation.ID, repo.deletedID)
	}
}

func TestDeleteReservationUseCase_NotFound(t *testing.T) {
	repo := &mockReservationRepo{findErr: errors.New("réservation non trouvée")}
	uc := NewDeleteReservationUseCase(repo)

	err := uc.Execute(context.Background(), "RES-123")
	if err == nil {
		t.Fatal("expected not found error, got nil")
	}
}
