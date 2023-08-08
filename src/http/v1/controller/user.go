package controller

import (
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

const userIdDB = 1

func Login(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login schema.Login
		var err error

		if err = validator.
			ParsLoginUser(c, &login); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		user := parser.SchemaToEntityLogin(login, userIdDB)

		if _, err = interactor.
			New(stor, logger).
			FindUser(c, user); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		t, err := generateToken(userIdDB)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": t})
	}
}

func UpdateUser(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := parser.GetUserID(c, userId)
		var updatePass schema.UpdateUser
		var err error

		if err = validator.
			ParsUpdateUser(
				c, &updatePass,
			); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		old, new := parser.
			SchemaToEntityUser(updatePass, userId)

		user, err := interactor.
			New(stor, logger).
			UpdateUser(c, old, new)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
