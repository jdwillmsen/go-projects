package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func envAccountSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAuthToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func envServiceSID() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}
