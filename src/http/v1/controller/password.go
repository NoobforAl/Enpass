package controller

import (
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/gin-gonic/gin"
)

func NewPass(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := getUserID(c)
		password := entity.Pass{UserID: userID}
		var err error

		if err = password.Pars(c); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		p, err := cachedPass.getPass(userID)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		if err = password.CreatePass(c, stor, p); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func AllPass(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := getUserID(c)
		var pass entity.Pass

		p, err := cachedPass.getPass(userID)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		passwords, err := pass.GetAllPassword(c, stor, p, true)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, passwords)
	}
}

func FindPass(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := getUserID(c)

		decrypt := c.Query("decrypt")
		passId, err := getParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		p, err := cachedPass.getPass(userID)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		password := entity.Pass{UserID: userID, PassID: uint(passId)}
		if err = password.FindPassword(c, stor, p, decrypt == "true"); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, password)
	}
}

func UpdatePass(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := getUserID(c)

		var pass entity.Pass
		var err error

		if err = pass.Pars(c); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		p, err := cachedPass.getPass(userId)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		if err = pass.UpdatePass(c, stor, p); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, pass)
	}
}

func DeletePassWord(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := getUserID(c)

		passId, err := getParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		password := entity.Pass{PassID: uint(passId), UserID: userID}
		if err = password.DeletePass(c, stor); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
	}
}
