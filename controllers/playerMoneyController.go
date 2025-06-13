package controllers

import (
	"EndioriteAPI/database"
	"EndioriteAPI/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPlayersMoney(c *gin.Context) {

	rows, err := database.DB.Query(`SELECT player, money FROM Economy`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database for all players' money"})
		return
	}
	defer rows.Close()

	playersMoney := make(map[string]float64)

	for rows.Next() {
		var player string
		var money float64

		if err := rows.Scan(&player, &money); err != nil {
			continue
		}

		playersMoney[player] = money
	}

	c.JSON(http.StatusOK, playersMoney)
}

func GetPlayerMoney(c *gin.Context) {
	username := c.Param("username")

	rows, err := database.DB.Query(`SELECT money FROM Economy WHERE player = ?`, username)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query database for username money.")
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}
	defer rows.Close()

	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found in the database."})
		return
	}

	var money float64
	if err := rows.Scan(&money); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning username money from database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"money": money})
}

func GetPlayersMoneyTop(c *gin.Context) {

	order, limit, offset, ok := utils.ParseTopParams(c)
	if !ok {
		return
	}

	query := fmt.Sprintf(`SELECT player, money FROM Economy ORDER BY money %s LIMIT ? OFFSET ?`, order)
	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving players' top money"})
		return
	}
	defer rows.Close()

	type PlayerMoney struct {
		Username string  `json:"username"`
		Money    float64 `json:"money"`
	}

	var playersMoneyTop []PlayerMoney

	for rows.Next() {
		var p PlayerMoney
		if err := rows.Scan(&p.Username, &p.Money); err != nil {
			continue
		}
		playersMoneyTop = append(playersMoneyTop, p)
	}

	c.JSON(http.StatusOK, playersMoneyTop)
}
