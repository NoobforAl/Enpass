package controller

import (
	"net/http"

	"github.com/NoobforAl/Enpass/Db"
	model "github.com/NoobforAl/Enpass/Model"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func NewPass(c *gin.Context) {
	userID := getUserID(c)

	var pass schema.Pass
	var err error

	if err = pass.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	err = Db.Get(&model.Service{ID: pass.ServiceID})
	if err != nil {
		errorHandling(c, err)
		return
	}

	values := model.Values{
		UserName: model.Value(pass.UserName),
		Password: model.Value(pass.Password),
		Note:     model.Value(pass.Note),
	}

	p, err := cachedPass.getPass(userID)
	if err != nil {
		errorHandling(c, err)
		return
	}

	if err = values.EncryptValues(p); err != nil {
		errorHandling(c, err)
		return
	}

	newPass := model.SavedPassword{
		UserPassID: uint(userID),
		ServiceID:  pass.ServiceID,
		Values:     values,
	}

	if err = Db.Insert(&newPass); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, pass)
}

func FindPass(c *gin.Context) {
	userID := getUserID(c)

	decrypt := c.Query("decrypt")
	passId, err := getParmInt(c, "id")
	if err != nil {
		errorHandling(c, err)
		return
	}

	pass := model.SavedPassword{ID: uint(passId), UserPassID: uint(userID)}
	if err = Db.Get(&pass); err != nil {
		errorHandling(c, err)
		return
	}

	if decrypt == "true" {
		p, err := cachedPass.getPass(userID)
		if err != nil {
			errorHandling(c, err)
			return
		}

		if err = pass.Values.DecryptValues(p); err != nil {
			errorHandling(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"UserName": pass.UserName,
		"Password": pass.Password,
		"Note":     pass.Note,
	})
}

func UpdatePass(c *gin.Context) {
	userId := getUserID(c)

	var pass schema.UpdatePass
	var err error

	if err = pass.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	values := model.Values{
		UserName: model.Value(pass.UserName),
		Password: model.Value(pass.Password),
		Note:     model.Value(pass.Note),
	}

	p, err := cachedPass.getPass(userId)
	if err != nil {
		errorHandling(c, err)
		return
	}

	if err = values.EncryptValues(p); err != nil {
		errorHandling(c, err)
		return
	}

	updatePass := model.SavedPassword{
		ID:         pass.PassID,
		UserPassID: uint(userId),
		ServiceID:  pass.ServiceID,
	}

	if err = Db.Get(&updatePass); err != nil {
		errorHandling(c, err)
		return
	}

	updatePass.Values = values
	if err = Db.Update(&updatePass); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, pass)
}

func DeletePassWord(c *gin.Context) {
	userID := getUserID(c)

	passId, err := getParmInt(c, "id")
	if err != nil {
		errorHandling(c, err)
		return
	}

	if err = Db.Delete(&model.SavedPassword{
		ID:         uint(passId),
		UserPassID: uint(userID),
	}); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
}
