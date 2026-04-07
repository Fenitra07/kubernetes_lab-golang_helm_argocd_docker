package entities

// UserID représente un identifiant unique d'utilisateur
type UserID int

// User représente un utilisateur dans le domaine métier
type User struct {
	ID       UserID
	Email    string
	Password string
}

// NewUser crée un nouvel utilisateur
func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

// IsValidCredentials vérifie si les identifiants sont valides
func (u *User) IsValidCredentials(email, password string) bool {
	return u.Email == email && u.Password == password
}