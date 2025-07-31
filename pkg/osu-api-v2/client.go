package osuapiv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL  = "https://osu.ppy.sh/api/v2"
	OauthURL = "https://osu.ppy.sh/oauth/token"
)

type OsuApiClient struct {
	ClientId     string
	ClientSecret string
	accessToken  string
}

type OauthRequestData struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type OauthResponseData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GetClientToken(client OsuApiClient) OauthResponseData {
	payload := OauthRequestData{
		ClientId:     client.ClientId,
		ClientSecret: client.ClientSecret,
		GrantType:    "client_credentials",
		Scope:        "public",
	}

	requestData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", OauthURL, bytes.NewBuffer(requestData))
	if err != nil {
		panic("Failed to create request: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic("Failed to get client token: " + err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var oauthData OauthResponseData

	err = json.Unmarshal(body, &oauthData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		panic(err)
	}

	return oauthData
}
