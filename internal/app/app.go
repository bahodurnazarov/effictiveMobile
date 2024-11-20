package app

import (
	"github.com/bahodurnazarov/effictiveMobile/internal/app/handler"
	"github.com/bahodurnazarov/effictiveMobile/internal/app/repository"
	"github.com/bahodurnazarov/effictiveMobile/internal/app/service"
	"github.com/bahodurnazarov/effictiveMobile/internal/router"
	"github.com/bahodurnazarov/effictiveMobile/pkg/db"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"github.com/gin-gonic/gin"
	"net"
)

type App struct {
	r *repository.Repository
	s *service.Service
	h *handler.Handler
	g *gin.Engine
}

func New(cfg *models.Settings) *App {
	app := &App{}
	err := utils.Logger(cfg)
	if err != nil {
		logger.DebugLogger.Println("Failed to initialize logger: %v", err)
	}

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		logger.DebugLogger.Println(err)
		return nil
	}

	app.r = repository.New(connDB)
	app.s = service.New(app.r)
	app.h = handler.New(app.s)
	app.g = router.Init(app.h)

	return app
}

func (a *App) Run(cfg *models.Settings) error {
	return a.g.Run(net.JoinHostPort(cfg.App.AppHost, cfg.App.PortRun))
}
