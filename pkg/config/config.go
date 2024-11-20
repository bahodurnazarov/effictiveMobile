package config

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Cnfg *models.Settings

func Init() (*models.Settings, error) {
	// Set the default environment to "prod"
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod" // Default to "prod" if not set
	}

	// Load environment-specific .env file
	envFile := ".env." + env
	viper.SetConfigFile(envFile) // Set the specific .env file directly
	viper.SetConfigType("env")   // Treat as .env format
	viper.AddConfigPath(".")     // Look for the .env file in the project root

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error reading environment config file (%s): %v", envFile, err)
	}

	// Automatically read system environment variables (overrides .env)
	viper.AutomaticEnv()

	// Unmarshal into the Settings struct
	var cfg models.Settings
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
		return nil, err
	}

	// Assign to the global variable
	Cnfg = &cfg
	return Cnfg, nil
}
