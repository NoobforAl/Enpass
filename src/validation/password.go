package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (_ validator) ParsPassword(
	c *gin.Context,
	pass *schema.Password,
) error {
	return c.BindJSON(pass)
}
