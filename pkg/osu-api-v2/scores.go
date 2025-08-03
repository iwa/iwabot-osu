package osuapiv2

type Score struct {
	Accuracy          float32         `json:"accuracy"`
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
	PP                float32         `json:"pp"`
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
	Beatmap           ScoreBeatmap    `json:"beatmap"`
	BeatmapSet        ScoreBeatmapSet `json:"beatmapset"`
}

type ScoreStatistics struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}

type ScoreBeatmap struct {
	BeatmapSetID     int     `json:"beatmapset_id"`
	DifficultyRating float32 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	Version          string  `json:"version"`
	TotalLength      int     `json:"total_length"`
	Accuracy         float32 `json:"accuracy"`
	Ar               float32 `json:"ar"`
	BPM              int     `json:"bpm"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
	CS               float32 `json:"cs"`
	Drain            float32 `json:"drain"`
	HitLength        int     `json:"hit_length"`
	LastUpdated      string  `json:"last_updated"`
	Passcount        int     `json:"passcount"`
	Playcount        int     `json:"playcount"`
	Ranked           int     `json:"ranked"`
	URL              string  `json:"url"`
	Checksum         string  `json:"checksum"`
}

type ScoreBeatmapSet struct {
	Artist         string `json:"artist"`
	ArtistUnicode  string `json:"artist_unicode"`
	Covers         Covers `json:"covers"`
	Creator        string `json:"creator"`
	FavouriteCount int    `json:"favourite_count"`
	ID             int    `json:"id"`
	Playcount      int    `json:"play_count"`
	Source         string `json:"source"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	TitleUnicode   string `json:"title_unicode"`
	UserId         int    `json:"user_id"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2x     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2x      string `json:"card@2x"`
	List        string `json:"list"`
	List2x      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2x string `json:"slimcover@2x"`
}
