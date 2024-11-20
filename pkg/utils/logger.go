package utils

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"log"
	"os"
	"path/filepath"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
)

func Logger(cfg *models.Settings) error {
	logFolder := "logs"
	if err := os.MkdirAll(logFolder, os.ModePerm); err != nil {
		return err
	}

	debugFilePath := filepath.Join(logFolder, cfg.App.DebugFile)
	infoFilePath := filepath.Join(logFolder, cfg.App.InfoFile)

	debugLog, err := os.OpenFile(debugFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	infoLog, err := os.OpenFile(infoFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	DebugLogger = log.New(debugLog, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(infoLog, "INFO: ", log.Ldate|log.Ltime)

	return nil
}
