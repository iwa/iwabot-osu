package utils

import "os"

func ImportClientInfoFromEnv() (string, string) {
	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		panic("[ERROR] No client id provided. Please set the CLIENT_ID environment variable.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		panic("[ERROR] No client secret provided. Please set the CLIENT_SECRET environment variable.")
	}

	return clientId, clientSecret
}
