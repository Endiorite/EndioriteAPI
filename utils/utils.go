package utils

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseTopParams(c *gin.Context) (order string, limit int, offset int, ok bool) {
	order = strings.ToUpper(c.DefaultQuery("order", "desc"))
	if order != "ASC" && order != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	p, err := strconv.Atoi(pageStr)
	if err != nil || p <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	l, err := strconv.Atoi(limitStr)
	if err != nil || l <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset = (p - 1) * l

	return order, l, offset, true
}
