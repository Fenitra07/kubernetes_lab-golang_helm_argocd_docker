package route

import (
	"net/http"
	"server/internal/config"
	"server/internal/controler"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRouter(apiAddress string) {
	cHandler := controler.NewHandler()
	r := gin.Default()
	r.Use(corsMiddleware())

	authGroup := r.Group(config.AuthPath)
	{
		authGroup.POST(config.Login, cHandler.Login)
	}

	reservationGroup := r.Group(config.ReservationPath)
	{
		reservationGroup.GET("", cHandler.ListReservations)
		reservationGroup.POST("", cHandler.CreateReservation)
		reservationGroup.PUT(config.ReservationID, cHandler.UpdateReservation)
		reservationGroup.DELETE(config.ReservationID, cHandler.DeleteReservation)
	}

	r.Run(apiAddress)
}

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

func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	if origin == "" {
		return false
	}

	for _, allowed := range allowedOrigins {
		if strings.EqualFold(origin, allowed) {
			return true
		}
	}

	return false
}
