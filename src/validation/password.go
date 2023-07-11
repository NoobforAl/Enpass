package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

type password interface {
	*schema.CreatePassword |
		*schema.UpdatePassword
}

func ParsPassword[T password](
	c *gin.Context,
	pass T,
) error {
	return c.BindJSON(pass)
}
