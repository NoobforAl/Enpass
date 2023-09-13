package controller

import (
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"

	"github.com/gin-gonic/gin"
)

const userIdDB = 1

func Login(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login schema.Login
		err := conf.Validation.ParsLoginUser(c, &login)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		user := parser.SchemaToEntityLogin(login, userIdDB)
		user, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).FindUser(c, user)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		t, err := generateToken(user.ID)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": t})
	}
}

func UpdateUser(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := parser.GetUserID(c, userId)
		var updatePass schema.UpdateUser
		err := conf.Validation.ParsUpdateUser(c, &updatePass)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		old, new := parser.
			SchemaToEntityUser(updatePass, userId)

		user, err := interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).UpdateUser(c, old, new)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
