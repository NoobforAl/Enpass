package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NoobforAl/Enpass/Db"
	model "github.com/NoobforAl/Enpass/Model"
	"github.com/gin-gonic/gin"
)

var (
	ErrUnexpectedSigning = errors.New("unexpected signing method:")
	ErrNotFoundPass      = errors.New("Not found Password saved")
)

func errorHandling(c *gin.Context, err error) {
	var status = http.StatusInternalServerError

	switch {
	case errors.Is(err, Db.ErrRecordNotFound):
		status = http.StatusNotFound

	case errors.Is(err, strconv.ErrSyntax):
		status = http.StatusBadRequest

	case errors.Is(err, model.ErrDecrypt):
	case errors.Is(err, model.ErrEncrypt):
	case errors.Is(err, model.ErrTextIsShort):
		status = http.StatusUnauthorized
	}

	c.JSON(status, gin.H{
		"detail": err.Error(),
	})
}
