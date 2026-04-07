package repository

import (
	"context"
	"server/internal/domain/entities"
)

// UserRepository définit le contrat pour l'accès aux données utilisateur
type UserRepository interface {
	// FindByCredentials trouve un utilisateur par email et mot de passe
	FindByCredentials(ctx context.Context, email, password string) (*entities.User, error)

	// FindByID trouve un utilisateur par son ID
	FindByID(ctx context.Context, id entities.UserID) (*entities.User, error)
}