package envs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "9010"),
		DBUser:     getEnv("DB_USER", "sz"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBAddress: fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/?authSource=%s",
			getEnv("DB_USER", "sz"),
			getEnv("DB_PASSWORD", "1234"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "27017"),
			getEnv("DB_NAME", "hrms"),
		),
		DBName: getEnv("DB_NAME", "hrms"),
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return fallback
	}
}
