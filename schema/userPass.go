package schema

import "github.com/gin-gonic/gin"

type Login struct {
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserPass struct {
	Login
}

func (l *Login) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, l)
}

func (u *UpdateUserPass) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, u)
}
