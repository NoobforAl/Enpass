package parser

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// userId
func GetUserID(
	c *gin.Context,
	key string,
) uint {
	str := c.GetString(key)
	userID, _ := strconv.Atoi(str)
	return uint(userID)
}

func GetQueryBool(
	c *gin.Context,
	key string,
) bool {
	s := c.Query(key)
	return strings.
		EqualFold(s, "true")
}

func GetQueryInt(
	c *gin.Context,
	key string,
) (int, error) {
	s := c.Query(key)
	return strconv.Atoi(s)
}

func GetParmInt(
	c *gin.Context,
	key string,
) (int, error) {
	id := c.Param(key)
	return strconv.Atoi(id)
}
