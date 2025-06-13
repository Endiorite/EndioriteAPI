package controllers

import (
	"EndioriteAPI/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPlayerCosmeticsList(c *gin.Context) {
	getPlayerCosmeticsList(c, "cosmetics")
}

func GetPlayerEquippedCosmeticsList(c *gin.Context) {
	getPlayerCosmeticsList(c, "equippedCosmetics")
}

func getPlayerCosmeticsList(c *gin.Context, listType string) {
	username := c.Param("username")

	query := fmt.Sprintf(`SELECT %s FROM cosmetics WHERE player = ?`, listType)
	rows, err := database.DB.Query(query, username)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query database for username %s list.", listType)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}
	defer rows.Close()

	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username not found in the database."})
		return
	}

	var cosmeticsListEncoded string
	if err := rows.Scan(&cosmeticsListEncoded); err != nil {
		errStr := fmt.Sprintf("Error scanning username %s from database.", listType)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}

	var cosmeticsList map[string]interface{}
	jsonErr := json.Unmarshal([]byte(cosmeticsListEncoded), &cosmeticsList)
	if jsonErr != nil {
		fmt.Printf("Error while decoding username %s cosmetics list: %v\n", username, jsonErr)
	}

	c.JSON(http.StatusOK, cosmeticsList)
}
