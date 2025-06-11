package handlers

import (
	"context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
	"ya41-56/internal/shared/response"
)

func PingHandler(dbConn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		dbInstance, err := dbConn.DB()
		if err != nil {
			log.Printf("failed to get DB instance: %v", err)
			http.Error(w, "unhealthy", http.StatusServiceUnavailable)
			return
		}

		if err := dbInstance.PingContext(ctx); err != nil {
			log.Printf("DB ping failed: %v", err)
			http.Error(w, "unhealthy", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		response.JSON(w, http.StatusOK, "pong")
	}
}
