package controler

import (
	"fmt"
	"net/http"
	"server/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *livraisonHandler) Login(c *gin.Context) {
	fmt.Println("Hello!")
	var loginData entity.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		fmt.Printf("Error parsing request body: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid login data",
		})
		return
	}

	if loginData.Login == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Login and password are required",
		})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Database connection not initialized",
		})
		return
	}

	var user entity.LoginRequest
	if err := h.db.Where("mail = ? AND motdepasse = ?", loginData.Login, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid login credentials",
		})
		return
	}

	response := entity.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login successful",
		"data":    response,
	})
}
