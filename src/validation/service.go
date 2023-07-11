package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

type service interface {
	*schema.CreateService |
		*schema.UpdateService
}

func ParsService[T service](
	c *gin.Context,
	pass T,
) error {
	return c.BindJSON(pass)
}
