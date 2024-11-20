package utils

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Cnfg *models.Settings

func Init() (*models.Settings, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	envFile := ".env." + env
	viper.SetConfigFile(envFile)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error reading file (%s): %v", envFile, err)
	}

	viper.AutomaticEnv()

	var cfg models.Settings
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode: %v", err)
		return nil, err
	}

	Cnfg = &cfg
	return Cnfg, nil
}
