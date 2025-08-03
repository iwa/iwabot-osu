package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/iwa/iwabot-osu/internal/discord/handlers"
	"github.com/iwa/iwabot-osu/internal/utils"
	osuapiv2 "github.com/iwa/iwabot-osu/pkg/osu-api-v2"
)

var (
	osuClientId     string
	osuClientSecret string
	discordToken    string
	devGuildId      string

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "pong!",
			Contexts: (*[]discordgo.InteractionContextType)(&[]discordgo.InteractionContextType{
				discordgo.InteractionContextGuild, discordgo.InteractionContextBotDM, discordgo.InteractionContextPrivateChannel,
			}),
			IntegrationTypes: (*[]discordgo.ApplicationIntegrationType)(&[]discordgo.ApplicationIntegrationType{
				discordgo.ApplicationIntegrationGuildInstall, discordgo.ApplicationIntegrationUserInstall,
			}),
		},
		{
			Name:        "rs",
			Description: "recent osu scores",
			Contexts: (*[]discordgo.InteractionContextType)(&[]discordgo.InteractionContextType{
				discordgo.InteractionContextGuild, discordgo.InteractionContextBotDM, discordgo.InteractionContextPrivateChannel,
			}),
			IntegrationTypes: (*[]discordgo.ApplicationIntegrationType)(&[]discordgo.ApplicationIntegrationType{
				discordgo.ApplicationIntegrationGuildInstall, discordgo.ApplicationIntegrationUserInstall,
			}),
		},
	}
)

func init() {
	osuClientId, osuClientSecret = utils.ImportOsuClientInfoFromEnv()
	discordToken, devGuildId = utils.ImportDiscordClientInfoFromEnv()

	if osuClientId == "" || osuClientSecret == "" {
		panic("OSU Client ID and Secret must be set in environment variables")
	}

	if discordToken == "" {
		panic("Discord token must be set in environment variables")
	}
}

func main() {
	log.Println("Starting iwabot-osu...")

	osuClient := utils.New(osuClientId, osuClientSecret)

	oauthData := osuapiv2.GetClientToken(*osuClient)
	if oauthData.AccessToken == "" {
		panic("Failed to retrieve access token")
	}

	osuClient.AccessToken = oauthData.AccessToken

	discordClient, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		panic("Failed to create Discord client: " + err.Error())
	}

	discordClient.AddHandler(messageCreate)
	discordClient.AddHandler(handlers.InteractionHandler)

	discordClient.Identify.Presence.Status = string(discordgo.StatusIdle)
	discordClient.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	err = discordClient.Open()
	if err != nil {
		panic("Failed to open Discord session: " + err.Error())
	}
	defer discordClient.Close()

	log.Println("Adding commands...")
	_, errcmds := discordClient.ApplicationCommandBulkOverwrite(discordClient.State.User.ID, devGuildId, commands)
	if errcmds != nil {
		log.Panicf("Cannot add commands: %v", errcmds)
	}

	log.Println("Bot is running. Press CTRL+C to exit.")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s

	// scores, err := osuClient.GetUserRecentScores(8423138, 1, "osu")
	// if err != nil {
	// 	panic("Failed to get user recent scores: " + err.Error())
	// }

	// for _, score := range scores {
	// 	println(" --- Score Details ---")
	// 	println("Score ID:", score.ID)
	// 	println("Beatmap Title:", score.BeatmapSet.Title)
	// 	println("Beatmap Diff:", score.Beatmap.Version)
	// 	println("User ID:", score.UserId)
	// 	println("PP:", score.PP)
	// 	println("Rank:", score.Rank)
	// 	println("Accuracy:", score.Accuracy)
	// 	println("Total Score:", score.TotalScore)
	// 	println("Max Combo:", score.MaxCombo)
	// 	println("Mods:", strings.Join(score.Mods, " "))
	// }
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
