package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (v validator) ParsUpdateUser(
	ctx *gin.Context,
	user *schema.UpdateUser,
) error {
	v.log.Debug("Pars Update User")
	return ctx.BindJSON(user)
}

func (v validator) ParsLoginUser(
	ctx *gin.Context,
	user *schema.Login,
) error {
	v.log.Debug("Pars Login User")
	return ctx.BindJSON(user)
}
