package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/application/dtos"
	"server/internal/application/usecases"
	"server/internal/domain/entities"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockReservationRepoController struct {
	saved     *entities.Reservation
	updated   *entities.Reservation
	deletedID entities.ReservationID
	findAll  []*entities.Reservation
	findByID *entities.Reservation
	saveErr  error
	findErr  error
	updateErr error
	deleteErr error
}

func (m *mockReservationRepoController) Save(ctx context.Context, reservation *entities.Reservation) error {
	if m.saveErr != nil {
		return m.saveErr
	}
	m.saved = reservation
	return nil
}

func (m *mockReservationRepoController) FindByID(ctx context.Context, id entities.ReservationID) (*entities.Reservation, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	return m.findByID, nil
}

func (m *mockReservationRepoController) FindAll(ctx context.Context) ([]*entities.Reservation, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	return m.findAll, nil
}

func (m *mockReservationRepoController) Update(ctx context.Context, reservation *entities.Reservation) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	m.updated = reservation
	return nil
}

func (m *mockReservationRepoController) Delete(ctx context.Context, id entities.ReservationID) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	m.deletedID = id
	return nil
}

type mockUserRepoController struct {
	user *entities.User
	err  error
}

func (m *mockUserRepoController) FindByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.user == nil || m.user.Email != email || m.user.Password != password {
		return nil, errors.New("identifiants invalides")
	}
	return m.user, nil
}

func (m *mockUserRepoController) FindByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.user, nil
}

func TestReservationController_CreateReservation_Success(t *testing.T) {
	repo := &mockReservationRepoController{}
	createUC := usecases.NewCreateReservationUseCase(repo)
	listUC := usecases.NewListReservationsUseCase(repo)
	updateUC := usecases.NewUpdateReservationUseCase(repo)
	deleteUC := usecases.NewDeleteReservationUseCase(repo)
	ctrl := NewReservationController(createUC, listUC, updateUC, deleteUC)

	gin.SetMode(gin.TestMode)
	reqBody, _ := json.Marshal(dtos.CreateReservationRequest{Nom: "Jean", Prenom: "Dupont", DepartVille: "TNR", ArriveeVille: "NBO", DateDepart: "2026-05-01", Motif: "Études", Classe: "Economique", Statut: "Confirmé"})
	req := httptest.NewRequest(http.MethodPost, "/reservations", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.CreateReservation(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}

func TestReservationController_UpdateReservation_NotFound(t *testing.T) {
	repo := &mockReservationRepoController{findErr: errors.New("réservation non trouvée")}
	createUC := usecases.NewCreateReservationUseCase(repo)
	listUC := usecases.NewListReservationsUseCase(repo)
	updateUC := usecases.NewUpdateReservationUseCase(repo)
	deleteUC := usecases.NewDeleteReservationUseCase(repo)
	ctrl := NewReservationController(createUC, listUC, updateUC, deleteUC)

	gin.SetMode(gin.TestMode)
	reqBody, _ := json.Marshal(dtos.UpdateReservationRequest{Motif: "Affaires"})
	req := httptest.NewRequest(http.MethodPut, "/reservations/RES-1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "RES-1"}}

	ctrl.UpdateReservation(c)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}

func TestReservationController_DeleteReservation_Success(t *testing.T) {
	repo := &mockReservationRepoController{findByID: entities.NewReservation("Jean", "Dupont", "TNR", "NBO", "Études", "Economique", "Confirmé")}
	createUC := usecases.NewCreateReservationUseCase(repo)
	listUC := usecases.NewListReservationsUseCase(repo)
	updateUC := usecases.NewUpdateReservationUseCase(repo)
	deleteUC := usecases.NewDeleteReservationUseCase(repo)
	ctrl := NewReservationController(createUC, listUC, updateUC, deleteUC)

	gin.SetMode(gin.TestMode)
	req := httptest.NewRequest(http.MethodDelete, "/reservations/RES-1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "RES-1"}}

	ctrl.DeleteReservation(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestAuthController_Login_Success(t *testing.T) {
	userRepo := &mockUserRepoController{user: &entities.User{Email: "test@example.com", Password: "secret"}}
	loginUC := usecases.NewLoginUseCase(userRepo)
	ctrl := NewAuthController(loginUC)

	gin.SetMode(gin.TestMode)
	reqBody, _ := json.Marshal(dtos.LoginRequest{Login: "test@example.com", Password: "secret"})
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Login(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	if body["mail"] != "test@example.com" {
		t.Fatalf("expected mail test@example.com, got %v", body["mail"])
	}
}

func TestAuthController_Login_InvalidPayload(t *testing.T) {
	userRepo := &mockUserRepoController{}
	loginUC := usecases.NewLoginUseCase(userRepo)
	ctrl := NewAuthController(loginUC)

	gin.SetMode(gin.TestMode)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString("{invalid json}"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Login(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}
