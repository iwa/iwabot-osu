package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iwa/iwabot-osu/internal/utils"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandName := i.ApplicationCommandData().Name
	if commandName == "ping" {
		ping(s, i)
		return
	}
	if commandName == "rs" {
		rs(s, i)
		return
	}
}

func ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong!",
		},
	})
}

func rs(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var user *discordgo.User
	if i.User != nil {
		user = i.User
	} else if i.Member != nil {
		user = i.Member.User
	}

	if user == nil {
		fmt.Println("Could not determine user")
		return
	}

	if user.ID != "125325519054045184" {
		return
	}

	const userID = "8423138"

	osuClient := utils.GetOsuClient()

	scores, err := osuClient.GetUserRecentScores(8423138, 1, "osu")
	if err != nil {
		panic("Failed to get user recent scores: " + err.Error())
	}

	if len(scores) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No recent scores found for user.",
			},
		})
		return
	}

	var component discordgo.MessageEmbed
	for _, score := range scores {
		component.Title = fmt.Sprintf("%s - %s [%s] [%.2f*]", score.BeatmapSet.Artist, score.BeatmapSet.Title, score.Beatmap.Version, float64(score.Beatmap.DifficultyRating))
		component.Description = fmt.Sprintf("%s - **%2.f** - **%2.fpp**/maxpp\n%d 300 - %d 100 - %d 50 - %dx\n%d • **%d**/%d\n`%d` • %dbpm\n%s", score.Rank, float64(score.Accuracy), float64(score.PP), score.Statistics.Count300, score.Statistics.Count100, score.Statistics.Count50, score.Statistics.CountMiss, score.TotalScore, score.MaxCombo, score.Beatmap.CountCircles+score.Beatmap.CountSliders+score.Beatmap.CountSpinners, score.Beatmap.TotalLength, score.Beatmap.BPM, strings.Join(score.Mods, " "))
		component.Color = 0x39CC69 // Green color for success
		component.Author = &discordgo.MessageEmbedAuthor{
			Name:    score.User.Username,
			IconURL: score.User.AvatarURL,
			URL:     fmt.Sprintf("https://osu.ppy.sh/users/%d", score.UserId),
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&component},
		},
	})

	log.Println("Command 'rs' executed by user:", user.GlobalName)
}
