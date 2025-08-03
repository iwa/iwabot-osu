package utils

import (
	"log"
	"os"
)

func ImportOsuClientInfoFromEnv() (string, string) {
	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		log.Panicln("[ERROR] No client id provided. Please set the CLIENT_ID environment variable.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Panicln("[ERROR] No client secret provided. Please set the CLIENT_SECRET environment variable.")
	}

	return clientId, clientSecret
}
