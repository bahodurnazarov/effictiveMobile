package app

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/db"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
)

type App struct {
}

func New(cfg *models.Settings) {
	err := utils.Logger(cfg)
	if err != nil {
		logger.DebugLogger.Println("Failed to initialize logger: %v", err)
	}
	logger.InfoLogger.Println("This is an info message.")

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		logger.DebugLogger.Println(err)
		return
	}
	logger.InfoLogger.Println(connDB)

}
