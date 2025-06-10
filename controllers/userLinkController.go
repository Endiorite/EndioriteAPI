package controllers

import (
	"fmt"

	"EndioriteAPI/database"
	"EndioriteAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserLink(c *gin.Context) {
	userId := c.Param("userId")

	rows, err := database.DB.Query(`SELECT username, linked FROM discord_link WHERE user_id = ?`, userId)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database for user link status."})
		return
	}
	defer rows.Close()

	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"isLinked": false, "error": "User not found in the database."})
		return
	}

	var username string
	var linked bool
	if err := rows.Scan(&username, &linked); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isLinked": false, "error": "Error scanning user data from database."})
		return
	}

	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isLinked": false, "error": "Error during row iteration."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"isLinked": linked, "username": username})
}

func LinkUser(c *gin.Context) {
	var reqBody models.UserLink
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("Missing or invalid data : %v", err.Error())})
		return
	}

	_, err = database.DB.Exec(
		`INSERT INTO discord_link (user_id, username, code) VALUES (?, ?, ?)`,
		reqBody.UserId,
		reqBody.Username,
		reqBody.Code,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to add information to the database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Link created successfully."})
}
