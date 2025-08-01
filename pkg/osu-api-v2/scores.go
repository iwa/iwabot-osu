package osuapiv2

type Score struct {
	Accuracy          float64         `json:"accuracy"`
	BeatmapId         int             `json:"beatmap_id"`
	ClassicTotalScore int             `json:"classic_total_score"`
	EndedAt           string          `json:"ended_at"`
	HasReplay         bool            `json:"has_replay"`
	ID                int             `json:"id"`
	IsPerfectCombo    bool            `json:"is_perfect_combo"`
	MaxCombo          int             `json:"max_combo"`
	Mods              []string        `json:"mods"`
	Passed            bool            `json:"passed"`
	PlaylistItemId    int             `json:"playlist_item_id"`
	PP                float64         `json:"pp"`
	Preserve          bool            `json:"preserve"`
	Processed         bool            `json:"processed"`
	Rank              string          `json:"rank"`
	Ranked            bool            `json:"ranked"`
	RoomId            int             `json:"room_id"`
	RulesetId         int             `json:"ruleset_id"`
	StartedAt         string          `json:"started_at"`
	Statistics        ScoreStatistics `json:"statistics"`
	TotalScore        int             `json:"total_score"`
	Type              string          `json:"type"`
	UserId            int             `json:"user_id"`
}

type ScoreStatistics struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}
