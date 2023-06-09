package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) uint {
	userID, _ := strconv.Atoi(c.GetString("userId"))
	return uint(userID)
}

func getQueryInt(c *gin.Context, key string) (int, error) {
	s := c.Query(key)
	return strconv.Atoi(s)
}

func getParmInt(c *gin.Context, key string) (int, error) {
	id := c.Param(key)
	return strconv.Atoi(id)
}
