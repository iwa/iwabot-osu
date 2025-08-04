package osuapiv2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *OsuApiClient) GetUserRecentScores(userId int, limit int, mode string) ([]Score, error) {
	url := fmt.Sprintf("%s/users/%d/scores/recent?include_fails=1&limit=%d", BaseURL, userId, limit)

	if mode != "" {
		url += fmt.Sprintf("&mode=%s", mode)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-version", "20240529")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user recent scores: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var scores []Score
	if err := json.NewDecoder(resp.Body).Decode(&scores); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return scores, nil
}
