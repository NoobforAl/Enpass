package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (v validator) ParsService(
	c *gin.Context,
	pass *schema.Service,
) error {
	v.log.Debug("Pars Service")
	return c.BindJSON(pass)
}
