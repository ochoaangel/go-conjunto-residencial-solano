package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	// check if prod
	prod := os.Getenv("PROD")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error cargando el archivo .env..")
		}
	}

	return nil
}
