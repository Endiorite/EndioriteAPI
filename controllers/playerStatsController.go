package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"EndioriteAPI/database"
	"EndioriteAPI/models"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
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

	order := c.DefaultQuery("order", "desc")
	if order != "asc" && order != "desc" && order != "ASC" && order != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter"})
		return
	}

	page := c.DefaultQuery("page", "1")
	p, err := strconv.Atoi(page)
	if err != nil || p <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit := c.DefaultQuery("limit", "10")
	l, err := strconv.Atoi(limit)
	if err != nil || l <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
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
		c.JSON(400, gin.H{"error": "Invalid topType"})
		return
	}

	offset := (p - 1) * l

	query := fmt.Sprintf(`
		SELECT username, %s
		FROM remstats
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, topType, topType, order)

	rows, err := database.DB.Query(query, l, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving player stats top"})
		return
	}
	defer rows.Close()

	type PlayerStat struct {
		Username string      `json:"username"`
		Stat     interface{} `json:"stat"`
	}

	var results []PlayerStat

	for rows.Next() {
		var username string
		var stat interface{}

		err := rows.Scan(&username, &stat)
		if err != nil {
			continue
		}

		results = append(results, PlayerStat{
			Username: username,
			Stat:     stat,
		})
	}

	c.JSON(http.StatusOK, results)
}
