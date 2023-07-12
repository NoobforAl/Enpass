package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (_ validator) ParsService(
	c *gin.Context,
	pass schema.Service,
) error {
	return c.BindJSON(pass)
}
