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

func ImportDiscordClientInfoFromEnv() (string, string) {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Panicln("[ERROR] No client id provided. Please set the CLIENT_ID environment variable.")
	}

	devGuild := os.Getenv("DEV_GUILD_ID")
	if devGuild == "" {
		log.Println("[INFO] No dev guild id provided, running in public mode.")
	} else {
		log.Printf("[INFO] Running in DEV mode with guild ID: %s", devGuild)
	}

	return token, devGuild
}
