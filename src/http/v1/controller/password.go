package controller

import (
	"fmt"
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

func NewPass(
	stor contract.Store,
	validator contract.Validation,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pass schema.Password
		var err error

		if err = validator.
			ParsPassword(c, &pass); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		userID := parser.GetUserID(c, userId)
		password := parser.SchemaToEntityPass(pass, 0, userID)

		password, err = interactor.
			New(stor).
			CreatePass(c, password, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func AllPass(
	stor contract.Store,
	validator contract.Validation,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parser.GetUserID(c, userId)
		decrypt := parser.GetQueryBool(c, "decrypt")

		passwords, err := interactor.
			New(stor).
			GetAllPassword(c, userID, decrypt)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, passwords)
	}
}

func FindPass(
	stor contract.Store,
	validator contract.Validation,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parser.GetUserID(c, userId)
		decrypt := parser.GetQueryBool(c, "decrypt")
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		pass := parser.SchemaToEntityPass(
			schema.Password{}, uint(id), userIdDB,
		)

		pass, err = interactor.
			New(stor).
			FindPassword(c, pass, userID, decrypt)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		fmt.Println(pass)
		c.JSON(http.StatusOK, pass)
	}
}

func UpdatePass(
	stor contract.Store,
	validator contract.Validation,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pass schema.Password
		var err error

		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		if err = validator.
			ParsPassword(c, &pass); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		userID := parser.GetUserID(c, userId)
		password := parser.SchemaToEntityPass(
			pass, uint(id), userID)

		password, err = interactor.
			New(stor).
			UpdatePass(c, password, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func DeletePass(
	stor contract.Store,
	validator contract.Validation,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		pass := parser.SchemaToEntityPass(
			schema.Password{}, uint(id), userIdDB,
		)

		pass, err = interactor.
			New(stor).
			DeletePass(c, pass)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, pass)
	}
}
