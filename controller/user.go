package controller

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/entity"
	env "github.com/NoobforAl/Enpass/loadEnv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = make([]byte, 32)

func init() {
	_, err := crand.Read(secretKey)
	if err != nil {
		panic(err)
	}
}

func GenRandomPass(c *gin.Context) {
	const (
		sizeAllChar = 88
		charSet     = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"0123456789!@#$%^&*()-_=+,.?/:;{}[]|~"
	)

	size, _ := getQueryInt(c, "size")
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

func generateToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(env.GetLifeTime()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Join(ErrUnexpectedSigning, fmt.Errorf("%v", token.Header["alg"]))
			}
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := fmt.Sprintf("%v", claims["id"])
		c.Set("userId", id)
		c.Next()
	}
}

func Login(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		var err error

		if err = user.Pars(c); err != nil {
			errorHandling(c, err)
			return
		}

		userid, err := user.FindUser(c, stor)
		if err != nil {
			errorHandling(c, err)
			return
		}

		t, err := generateToken(userid)
		if err != nil {
			errorHandling(c, err)
			return
		}

		if _, err = cachedPass.getPass(userid); err != nil {
			go cachedPass.deletePass(userid)
		}

		cachedPass.setPass(userid, user.Password)
		c.JSON(http.StatusOK, gin.H{"token": t})
	}
}

func UpdateUser(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := getUserID(c)
		var user entity.User
		var err error

		if err = user.Pars(c); err != nil {
			errorHandling(c, err)
			return
		}

		oldPass, err := cachedPass.getPass(userId)
		if err != nil {
			errorHandling(c, err)
			return
		}

		if oldPass != user.Password {
			errorHandling(c, ErrNotMatchPassword)
			return
		}

		if err = user.UpdateUser(c, stor); err != nil {
			errorHandling(c, err)
			return
		}

		cachedPass.setPass(userId, user.Password)
		c.JSON(http.StatusBadRequest, user)
	}
}
