package controllers

import (
	"EndioriteAPI/database"
	"EndioriteAPI/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserLink(c *gin.Context) {
	userId := c.Param("userId")

	rows, err := database.DB.Query(`SELECT username, linked FROM discord_link WHERE user_id = ?`, userId)
	if err != nil {
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

	result, err := database.DB.Exec(
		`INSERT INTO discord_link (user_id, username, code)
				VALUES (?, ?, ?)
				ON DUPLICATE KEY UPDATE
					username = IF(TIMESTAMPDIFF(SECOND, code_date, NOW()) > 300, ?, username),
					code     = IF(TIMESTAMPDIFF(SECOND, code_date, NOW()) > 300, ?, code),
					code_date = IF(TIMESTAMPDIFF(SECOND, code_date, NOW()) > 300, DEFAULT(code_date), code_date);
			`, reqBody.UserId, reqBody.Username, reqBody.Code, reqBody.Username, reqBody.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed attempt to insert or update the link."})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to get rows affected from insert operation."})
		return
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successfully updated user link."})
	} else {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "The user already has a code and it has not expired. Please wait 5 minutes."})
	}
}

func UnlinkUser(c *gin.Context) {
	userId := c.Param("userId")

	_, err := database.DB.Query(`DELETE FROM discord_link WHERE user_id = ?`, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete from the database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Link deleted successfully."})
}
