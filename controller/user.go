package controller

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/NoobforAl/Enpass/Db"
	model "github.com/NoobforAl/Enpass/Model"
	env "github.com/NoobforAl/Enpass/loadEnv"
	"github.com/NoobforAl/Enpass/schema"
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

func Login(c *gin.Context) {
	var loginVal schema.Login
	if err := loginVal.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	password := loginVal.Password
	data, err := Db.GetMany(&model.UserPass{})
	if err != nil {
		errorHandling(c, err)
		return
	}

	for _, v := range data {
		val, err := v.EnPass.DecryptValue(password)
		if err == nil && val.IsOkHash(password) {
			t, err := generateToken(v.ID)
			if err != nil {
				errorHandling(c, err)
				return
			}

			if _, err := cachedPass.getPass(v.ID); err != nil {
				go cachedPass.deletePass(v.ID)
			}

			cachedPass.setPass(v.ID, password)
			c.JSON(http.StatusOK, gin.H{"token": t})
			return
		}
	}

	errorHandling(c, Db.ErrRecordNotFound)
}

func UpdateUserPass(c *gin.Context) {
	userId := getUserID(c)

	var pass schema.UpdateUserPass
	var err error

	if err = pass.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	oldPass, err := cachedPass.getPass(userId)
	if err != nil {
		errorHandling(c, err)
		return
	}

	userPass := model.UserPass{
		ID:     uint(userId),
		EnPass: model.Value(pass.Password),
	}

	userPass.EnPass = userPass.EnPass.HashSha256()
	userPass.EnPass, err = userPass.EnPass.EncryptValue(pass.Password)
	if err != nil {
		errorHandling(c, err)
		return
	}

	if err = Db.Update(&userPass); err != nil {
		errorHandling(c, err)
		return
	}

	allPass, err := Db.GetMany(&model.SavedPassword{UserPassID: uint(userId)})
	if err != nil {
		userPass.EnPass = model.Value(oldPass)
		userPass.EnPass = userPass.EnPass.HashSha256()
		userPass.EnPass, _ = userPass.EnPass.EncryptValue(pass.Password)
		e := Db.Update(&userPass)
		if e != nil {
			err = errors.Join(err, e)
		}

		errorHandling(c, err)
		return
	}

	for i := range allPass {
		_ = allPass[i].Values.DecryptValues(oldPass)
		_ = allPass[i].Values.EncryptValues(pass.Password)
	}

	if err = Db.UpdateMany(allPass); err != nil {
		userPass.EnPass = model.Value(oldPass)
		userPass.EnPass = userPass.EnPass.HashSha256()
		userPass.EnPass, _ = userPass.EnPass.EncryptValue(pass.Password)
		e := Db.Update(&userPass)
		if e != nil {
			err = errors.Join(err, e)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	cachedPass.setPass(userId, pass.Password)
	c.JSON(http.StatusBadRequest, allPass)
}

func GenRandomPass(c *gin.Context) {
	const (
		sizeAllChar = 88
		charSet     = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"0123456789!@#$%^&*()-_=+,.?/:;{}[]|~"
	)

	size, _ := getQueryInt(c, "size")
	if size == 0 {
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

func generateToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(env.GetLifeTime()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
