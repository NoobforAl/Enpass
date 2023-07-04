package entity

import "github.com/gin-gonic/gin"

type schema interface {
	*Pass | *Service | *User
}

func parsJsonAndValidate[T schema](c *gin.Context, val T) error {
	return c.BindJSON(val)
}
