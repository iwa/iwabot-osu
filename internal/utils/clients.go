package utils

import osuapiv2 "github.com/iwa/iwabot-osu/pkg/osu-api-v2"

var (
	OsuClient osuapiv2.OsuApiClient
)

func New(clientId string, clientSecret string) *osuapiv2.OsuApiClient {
	OsuClient = osuapiv2.OsuApiClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	return &OsuClient
}

func GetOsuClient() *osuapiv2.OsuApiClient {
	return &OsuClient
}
