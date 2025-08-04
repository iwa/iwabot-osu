package osuapiv2

type Score struct {
	Accuracy          float64 `json:"accuracy"`
	BeatmapId         int     `json:"beatmap_id"`
	ClassicTotalScore int     `json:"classic_total_score"`
	EndedAt           string  `json:"ended_at"`
	HasReplay         bool    `json:"has_replay"`
	ID                int     `json:"id"`
	IsPerfectCombo    bool    `json:"is_perfect_combo"`
	MaxCombo          int     `json:"max_combo"`
	MaximumStatistics struct {
		Great         int `json:"great"`
		IgnoreHit     int `json:"ignore_hit"`
		LargeTickHit  int `json:"large_tick_hit"`
		SliderTailHit int `json:"slider_tail_hit"`
	} `json:"maximum_statistics"`
	Mods           []string        `json:"mods"`
	Passed         bool            `json:"passed"`
	PlaylistItemId int             `json:"playlist_item_id"`
	PP             float64         `json:"pp"`
	Preserve       bool            `json:"preserve"`
	Processed      bool            `json:"processed"`
	Rank           string          `json:"rank"`
	Ranked         bool            `json:"ranked"`
	RoomId         int             `json:"room_id"`
	RulesetId      int             `json:"ruleset_id"`
	StartedAt      string          `json:"started_at"`
	Statistics     ScoreStatistics `json:"statistics"`
	TotalScore     int             `json:"total_score"`
	Type           string          `json:"type"`
	UserId         int             `json:"user_id"`
	Beatmap        ScoreBeatmap    `json:"beatmap"`
	BeatmapSet     ScoreBeatmapSet `json:"beatmapset"`
	User           ScoreUser       `json:"user"`
}

type ScoreStatistics struct {
	CountMiss     int `json:"miss"`
	Count50       int `json:"meh"`
	Count100      int `json:"ok"`
	Count300      int `json:"great"`
	IgnoreHit     int `json:"ignore_hit"`
	LargeTickHit  int `json:"large_tick_hit"`
	SliderTailHit int `json:"slider_tail_hit"`
}

type ScoreBeatmap struct {
	BeatmapSetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	Version          string  `json:"version"`
	TotalLength      int     `json:"total_length"`
	Accuracy         float64 `json:"accuracy"`
	Ar               float64 `json:"ar"`
	BPM              int     `json:"bpm"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
	CS               float64 `json:"cs"`
	Drain            float64 `json:"drain"`
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

type ScoreUser struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	AvatarURL   string `json:"avatar_url"`
	CountryCode string `json:"country_code"`
}
