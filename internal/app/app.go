package app

import (
	"fmt"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"log"
)

type App struct {
}

func New(cfg *models.Settings) {
	err := utils.Logger(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	// Example log messages
	logger.DebugLogger.Println("This is a debug message.")
	logger.InfoLogger.Println("This is an info message.")

	fmt.Println("Application started successfully.")

}
