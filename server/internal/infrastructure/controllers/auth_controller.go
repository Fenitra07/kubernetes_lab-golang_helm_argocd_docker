package controllers

import (
	"net/http"
	"server/internal/application/dtos"
	"server/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

// AuthController gère les requêtes HTTP d'authentification
type AuthController struct {
	loginUseCase *usecases.LoginUseCase
}

// NewAuthController crée une nouvelle instance du contrôleur
func NewAuthController(loginUC *usecases.LoginUseCase) *AuthController {
	return &AuthController{
		loginUseCase: loginUC,
	}
}

// Login gère la requête POST /auth/login
func (ctrl *AuthController) Login(c *gin.Context) {
	var request dtos.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid login data",
		})
		return
	}

	response, err := ctrl.loginUseCase.Execute(c.Request.Context(), request)
	if err != nil {
		status := http.StatusUnauthorized
		if err.Error() == "login et mot de passe sont requis" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{
			"status":  status,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}