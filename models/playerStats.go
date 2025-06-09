package models

type PlayerStats struct {
	XUID           string `json:"xuid"`
	Username       string `json:"username"`
	Kills          int64  `json:"kills"`
	Deaths         int64  `json:"deaths"`
	KillStreak     int64  `json:"kill_streak"`
	BestKillStreak int64  `json:"best_kill_streak"`
	PlayingTime    int64  `json:"playing_time"`
}
