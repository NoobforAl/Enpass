package controller

import (
	"math/rand"
	"net/http"

	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/gin-gonic/gin"
)

func GenRandomPass(c *gin.Context) {
	const (
		sizeAllChar = 88
		charSet     = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"0123456789!@#$%^&*()-_=+,.?/:;{}[]|~"
	)

	size, _ := parser.GetQueryInt(c, "size")
	if size <= 0 || size > 1000 {
		size = 10
	}

	passWord := make([]byte, size)
	for i := 0; i < size; i++ {
		passWord[i] = charSet[rand.Intn(sizeAllChar)]
	}

	c.JSON(http.StatusOK, gin.H{
		"password": string(passWord),
	})
}
