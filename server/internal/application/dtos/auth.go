package dtos

// LoginRequest représente la requête de connexion
type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse représente la réponse de connexion
type LoginResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"mail"`
	Password string `json:"motdepasse"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}