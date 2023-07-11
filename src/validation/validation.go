package validation

import (
	"github.com/NoobforAl/Enpass/entity"
	"github.com/gin-gonic/gin"
)

type schema interface {
	*entity.User |
		*entity.Service |
		*entity.Password
}

func ParsJsonAndValidate[T schema](c *gin.Context, val T) error {
	return c.BindJSON(val)
}
