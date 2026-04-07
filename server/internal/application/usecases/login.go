package usecases

import (
	"context"
	"errors"
	"server/internal/application/dtos"
	"server/internal/interfaces/repository"
)

// LoginUseCase gère l'authentification des utilisateurs
type LoginUseCase struct {
	userRepo repository.UserRepository
}

// NewLoginUseCase crée une nouvelle instance du use case
func NewLoginUseCase(repo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepo: repo,
	}
}

// Execute exécute le use case d'authentification
func (uc *LoginUseCase) Execute(ctx context.Context, req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	// Validation des entrées
	if req.Login == "" || req.Password == "" {
		return nil, errors.New("login et mot de passe sont requis")
	}

	// Recherche de l'utilisateur via le repository
	user, err := uc.userRepo.FindByCredentials(ctx, req.Login, req.Password)
	if err != nil {
		return nil, errors.New("identifiants invalides")
	}

	// Conversion vers DTO de réponse
	response := &dtos.LoginResponse{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
		Status:   200,
		Message:  "Connexion réussie",
	}

	return response, nil
}