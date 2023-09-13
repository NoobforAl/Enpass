package controller

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"time"

	env "github.com/NoobforAl/Enpass/config_loader"
	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = make([]byte, 32)

const userId = "userId"

func init() {
	if tmp := env.GetSecretKey(); len(tmp) > 4 {
		secretKey = tmp
		return
	}

	_, err := rand.Read(secretKey)
	if err != nil {
		panic(err)
	}
}

func AuthMiddleware(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Join(errs.ErrUnexpectedSigning, fmt.Errorf("%v", token.Header["alg"]))
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
		c.Set(userId, id)
		c.Next()
	}
}

func generateToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(env.GetLifeTime()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
