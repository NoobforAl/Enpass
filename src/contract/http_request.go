package contract

import "github.com/gin-gonic/gin"

type httpReq interface {
	GetUserID(
		c *gin.Context,
		key string,
	) uint

	GetQueryInt(
		c *gin.Context,
		key string,
	) (int, error)

	GetParmInt(
		c *gin.Context,
		key string,
	) (int, error)
}
