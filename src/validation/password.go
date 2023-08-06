package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (v validator) ParsPassword(
	c *gin.Context,
	pass *schema.Password,
) error {
	v.log.Debug("Pars Password")
	return c.BindJSON(pass)
}
