package controller

import (
	"fmt"
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"

	"github.com/gin-gonic/gin"
)

func NewPass(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pass schema.Password
		err := conf.Validation.ParsPassword(c, &pass)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		userID := parser.GetUserID(c, userId)
		password := parser.SchemaToEntityPass(pass, 0, userID)

		password, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).CreatePass(c, password, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func AllPass(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parser.GetUserID(c, userId)
		decrypt := parser.GetQueryBool(c, "decrypt")

		passwords, err := interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).GetAllPassword(c, userID, decrypt)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, passwords)
	}
}

func FindPass(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parser.GetUserID(c, userId)
		//decrypt := parser.GetQueryBool(c, "decrypt")
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		pass := parser.SchemaToEntityPass(
			schema.Password{}, uint(id), userIdDB,
		)

		pass, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).FindPassword(c, pass, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		fmt.Println(pass)
		c.JSON(http.StatusOK, pass)
	}
}

func UpdatePass(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pass schema.Password

		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		err = conf.Validation.ParsPassword(c, &pass)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		userID := parser.GetUserID(c, userId)
		password := parser.SchemaToEntityPass(
			pass, uint(id), userID)

		password, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).UpdatePass(c, password, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func DeletePass(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parser.GetUserID(c, userId)
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		pass := parser.SchemaToEntityPass(
			schema.Password{}, uint(id), userIdDB,
		)

		pass, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).DeletePass(c, pass, userID)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, pass)
	}
}
