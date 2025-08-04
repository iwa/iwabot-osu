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
		// fmt.Printf("score: %+v\n", score)

		component.Title = fmt.Sprintf("%s - %s [%s] [%.2f*]", score.BeatmapSet.Artist, score.BeatmapSet.Title, score.Beatmap.Version, float64(score.Beatmap.DifficultyRating))
		component.URL = score.Beatmap.URL

		component.Description = fmt.Sprintf("**%s** ・ **%1.2f%%** ・ **%1.2fpp**/maxpp\n%d <:300:1402045799164022914> %d <:100:1402045789513191566> %d <:50:1402045777324413040> %d <:X2:1402045819758317750>\n%d ・ **%dx**/%dx\n`%d:%02d` ・ %dbpm\n%s",
			score.Rank,
			(score.Accuracy * 100),
			score.PP,
			score.Statistics.Count300,
			score.Statistics.Count100,
			score.Statistics.Count50,
			score.Statistics.CountMiss,
			score.TotalScore,
			score.MaxCombo,
			(score.MaximumStatistics.Great + score.MaximumStatistics.LargeTickHit + score.MaximumStatistics.SliderTailHit),
			(score.Beatmap.TotalLength / 60),
			(score.Beatmap.TotalLength % 60),
			score.Beatmap.BPM,
			strings.Join(score.Mods, " "))

		component.Color = 0x7b67f0
		component.Author = &discordgo.MessageEmbedAuthor{
			Name:    score.User.Username,
			IconURL: score.User.AvatarURL,
			URL:     fmt.Sprintf("https://osu.ppy.sh/users/%d", score.UserId),
		}
		component.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: score.BeatmapSet.Covers.List2x,
		}
		component.Footer = &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Mapped by %s", score.BeatmapSet.Creator),
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
