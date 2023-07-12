package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

type user interface {
	*schema.UpdateUser |
		*schema.GetUser
}

func ParsUser[T user](
	c *gin.Context,
	user T,
) error {
	return c.BindJSON(user)
}
