package routes

import (
	"net/http"
	"server/internal/config"
	"server/internal/interfaces/presenters"

	"github.com/gin-gonic/gin"
)

// Router gère la configuration des routes
type Router struct {
	reservationController presenters.ReservationPresenter
	authController        presenters.AuthPresenter
}

// NewRouter crée une nouvelle instance du routeur
func NewRouter(
	reservationCtrl presenters.ReservationPresenter,
	authCtrl presenters.AuthPresenter,
) *Router {
	return &Router{
		reservationController: reservationCtrl,
		authController:        authCtrl,
	}
}

// SetupRouter configure et retourne le routeur Gin
func (r *Router) SetupRouter(apiAddress string) {
	ginRouter := gin.Default()
	ginRouter.Use(corsMiddleware())

	authGroup := ginRouter.Group(config.AuthPath)
	{
		authGroup.POST(config.Login, r.authController.Login)
	}

	reservationGroup := ginRouter.Group(config.ReservationPath)
	{
		reservationGroup.GET("", r.reservationController.ListReservations)
		reservationGroup.POST("", r.reservationController.CreateReservation)
		reservationGroup.PUT(config.ReservationID, r.reservationController.UpdateReservation)
		reservationGroup.DELETE(config.ReservationID, r.reservationController.DeleteReservation)
	}

	ginRouter.Run(apiAddress)
}

// corsMiddleware gère les en-têtes CORS
func corsMiddleware() gin.HandlerFunc {
	allowedOrigins := append([]string{}, config.AppConfiguration.API.AllowOrigins...)
	allowedOrigins = append(allowedOrigins, "null")

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if isAllowedOrigin(origin, allowedOrigins) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
		}

		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json; charset=utf-8")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// isAllowedOrigin vérifie si l'origine est autorisée
func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}