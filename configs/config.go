package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	DBConn  string // FATAL if not existing
	AppUrl  string
	AppName string
	// Add more game-specific config here
}

// utilise the godotenv package so that my env file
// is loaded into the current session and then accessible
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No ENV file found to load into environment")
	} else {
		log.Println("Env Loaded successfully")
	}
}

// set the defaults required for the project to run
func Load() *Config {
	return &Config{
		Port:    GetEnv("APP_PORT", "5001"),
		DBConn:  GetEnv("DB_CONN", ""),
		AppUrl:  GetEnv("APP_ADDRESS", "http://localhost"),
		AppName: GetEnv("APP_NAME", "PenisFlaps"), // just to ensure im reading from .env
	}
}

// helper function to read the .env file and assign
// the values or uses the passed in fallback
func GetEnv(envKey, fallbackVal string) string {
	// check if there is an environment variable
	// set for this key and get the value.
	// else return the fallback value
	envValue, exists := os.LookupEnv(envKey)
	if !exists || envValue == "" {
		if fallbackVal == "" {
			log.Fatalf("Required ENV Variable %s not found and no fallback provided", envKey)
		}
		return fallbackVal
	}
	return envValue
}
