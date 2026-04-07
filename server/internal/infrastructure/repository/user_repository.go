package repository

import (
	"context"
	"errors"
	"server/internal/domain/entities"
	"server/internal/interfaces/repository"

	"gorm.io/gorm"
)

// UserModel représente le modèle GORM pour la persistance des utilisateurs
type UserModel struct {
	ID       int    `gorm:"column:id;primaryKey"`
	Email    string `gorm:"column:mail"`
	Password string `gorm:"column:motdepasse"`
}

func (UserModel) TableName() string {
	return "login"
}

// gormUserRepository implémente UserRepository avec GORM
type gormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository crée une nouvelle instance du repository GORM
func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &gormUserRepository{db: db}
}

// FindByCredentials trouve un utilisateur par email et mot de passe
func (r *gormUserRepository) FindByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	var model UserModel
	err := r.db.WithContext(ctx).Where("mail = ? AND motdepasse = ?", email, password).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("utilisateur non trouvé")
		}
		return nil, err
	}
	return toUserEntity(&model), nil
}

// FindByID trouve un utilisateur par son ID
func (r *gormUserRepository) FindByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	var model UserModel
	err := r.db.WithContext(ctx).Where("id = ?", int(id)).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("utilisateur non trouvé")
		}
		return nil, err
	}
	return toUserEntity(&model), nil
}

// toUserEntity convertit un modèle GORM vers l'entité domaine
func toUserEntity(m *UserModel) *entities.User {
	return &entities.User{
		ID:       entities.UserID(m.ID),
		Email:    m.Email,
		Password: m.Password,
	}
}