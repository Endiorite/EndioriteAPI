package controllers

import (
	"EndioriteAPI/database"
	"EndioriteAPI/models"
	"EndioriteAPI/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPlayersStats(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT xuid, username, kills, deaths, kill_streak, best_kill_streak, playing_time
		FROM remstats
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving player stats"})
		return
	}
	defer rows.Close()

	var PlayersStats []models.PlayerStats

	for rows.Next() {
		var PlayerStats models.PlayerStats
		err := rows.Scan(
			&PlayerStats.XUID,
			&PlayerStats.Username,
			&PlayerStats.Kills,
			&PlayerStats.Deaths,
			&PlayerStats.KillStreak,
			&PlayerStats.BestKillStreak,
			&PlayerStats.PlayingTime,
		)
		if err != nil {
			continue
		}
		PlayersStats = append(PlayersStats, PlayerStats)
	}

	c.JSON(http.StatusOK, PlayersStats)
}

func GetPlayerStats(c *gin.Context) {
	xuid := c.Query("xuid")
	username := c.Query("username")

	var row *sql.Row
	if xuid != "" {
		row = database.DB.QueryRow(`
			SELECT xuid, username, kills, deaths, kill_streak, best_kill_streak, playing_time
			FROM remstats WHERE xuid = ?
		`, xuid)
	} else if username != "" {
		row = database.DB.QueryRow(`
			SELECT xuid, username, kills, deaths, kill_streak, best_kill_streak, playing_time
			FROM remstats WHERE username = ?
		`, username)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing xuid or username parameter"})
		return
	}

	var playerStats models.PlayerStats
	err := row.Scan(
		&playerStats.XUID,
		&playerStats.Username,
		&playerStats.Kills,
		&playerStats.Deaths,
		&playerStats.KillStreak,
		&playerStats.BestKillStreak,
		&playerStats.PlayingTime,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PlayerMoney not found"})
		return
	}

	c.JSON(http.StatusOK, playerStats)
}

func GetPlayersStatsTopKills(c *gin.Context) {
	getPlayersStatsTop(c, "kills")
}

func GetPlayersStatsTopDeaths(c *gin.Context) {
	getPlayersStatsTop(c, "deaths")
}

func GetPlayersStatsTopKillStreak(c *gin.Context) {
	getPlayersStatsTop(c, "kill_streak")
}

func GetPlayersStatsTopBestKillStreak(c *gin.Context) {
	getPlayersStatsTop(c, "best_kill_streak")
}

func GetPlayersStatsTopPlayingTime(c *gin.Context) {
	getPlayersStatsTop(c, "playing_time")
}

func getPlayersStatsTop(c *gin.Context, topType string) {

	order, limit, offset, ok := utils.ParseTopParams(c)
	if !ok {
		return
	}

	validTopTypes := map[string]bool{
		"kills":            true,
		"deaths":           true,
		"kill_streak":      true,
		"best_kill_streak": true,
		"playing_time":     true,
	}

	if !validTopTypes[topType] {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid topType"})
		return
	}

	query := fmt.Sprintf(`SELECT username, %s FROM remstats ORDER BY %s %s LIMIT ? OFFSET ?`, topType, topType, order)
	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving player stats top"})
		return
	}
	defer rows.Close()

	type PlayerStat struct {
		Username string `json:"username"`
		Stat     int64  `json:"stat"`
	}

	var playersStatTop []PlayerStat

	for rows.Next() {
		var p PlayerStat
		if err := rows.Scan(&p.Username, &p.Stat); err != nil {
			continue
		}
		playersStatTop = append(playersStatTop, p)
	}

	c.JSON(http.StatusOK, playersStatTop)
}
