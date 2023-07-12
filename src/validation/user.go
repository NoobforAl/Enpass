package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func (_ validator) ParsUpdateUser(
	ctx *gin.Context,
	user *schema.UpdateUser,
) error {
	return ctx.BindJSON(user)
}

func (_ validator) ParsLoginUser(
	ctx *gin.Context,
	user *schema.Login,
) error {
	return ctx.BindJSON(user)
}
