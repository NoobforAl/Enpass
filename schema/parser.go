package schema

import "github.com/gin-gonic/gin"

type schema interface {
	*Pass | *UpdatePass |
		*Service | *UpdateService |
		*Login | *UpdateUserPass
}

func parsJsonAndValidate[T schema](c *gin.Context, val T) error {
	return c.BindJSON(val)
}
