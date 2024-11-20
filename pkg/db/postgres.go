package db

import (
	"database/sql"
	"fmt"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB(cfg *models.Settings) (*gorm.DB, error) {
	db, err := postgresDB(cfg)
	return db, err

}

func postgresDB(cfg *models.Settings) (*gorm.DB, error) {
	database := cfg.Postgres

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.PGHost, database.PGPort, database.PGUser, database.PGUser, database.PGName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sql.Open("postgres", dsn)
	if err != nil {
		logger.DebugLogger.Println(err)
		return nil, err
	}
	if err = db.AutoMigrate(); err != nil {
		logger.DebugLogger.Println(err.Error())
		return nil, err
	}
	DB = db

	logger.InfoLogger.Println("*** POSTGRES *** -> success connect to Database by user")
	return DB, nil
}
