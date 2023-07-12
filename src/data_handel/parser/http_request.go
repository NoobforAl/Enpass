package parser

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// userId
func (_ httpPars) GetUserID(
	c *gin.Context,
	key string,
) uint {
	str := c.GetString(key)
	userID, _ := strconv.Atoi(str)
	return uint(userID)
}

func (_ httpPars) GetQueryInt(
	c *gin.Context,
	key string,
) (int, error) {
	s := c.Query(key)
	return strconv.Atoi(s)
}

func (_ httpPars) GetParmInt(
	c *gin.Context,
	key string,
) (int, error) {
	id := c.Param(key)
	return strconv.Atoi(id)
}
