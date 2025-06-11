package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	httpServer "net/http"
	"ya41-56/cmd"
	"ya41-56/internal/gophermart/db"
	"ya41-56/internal/gophermart/http"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/gophermart/repositories"
	"ya41-56/internal/gophermart/services"
	"ya41-56/internal/shared/di"
	"ya41-56/internal/shared/logger"
)

func Run() {
	cfg := cmd.ParseFlags()
	logger.Init(cfg.ModeLogger)

	dbConn := db.InitPostgres(&db.InitPostgresConfig{
		DSN:             cfg.DatabaseDSN,
		IsFireMigration: true,
	})

	if dbConn == nil {
		logger.L().Fatal("failed to init database")
	}

	userRepo := repositories.NewGormRepository[models.User](dbConn)

	r := http.RegisterRoutes(&di.AppContainer{
		UserRepo: userRepo,
		Auth:     services.NewAuthService(userRepo), // тут должен быть сервис для авторизации, сейчас нет реализации
		Router:   chi.NewRouter(),
		Cfg:      cfg,
		Gorm:     dbConn,
	})

	logger.L().Info("starting HTTP server", zap.String("addr", cfg.Address))

	err := httpServer.ListenAndServe(cfg.Address, r)
	if err != nil {
		logger.L().Fatal("failed to start HTTP server", zap.Error(err))
	}
}
