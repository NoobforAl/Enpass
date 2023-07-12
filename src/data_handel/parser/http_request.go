package parser

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// userId
func (_ parser) GetUserID(
	c *gin.Context,
	key string,
) uint {
	str := c.GetString(key)
	userID, _ := strconv.Atoi(str)
	return uint(userID)
}

func (_ parser) GetQueryInt(
	c *gin.Context,
	key string,
) (int, error) {
	s := c.Query(key)
	return strconv.Atoi(s)
}

func (_ parser) GetParmInt(
	c *gin.Context,
	key string,
) (int, error) {
	id := c.Param(key)
	return strconv.Atoi(id)
}
