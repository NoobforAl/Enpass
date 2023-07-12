package main

import (
	env "github.com/NoobforAl/Enpass/config_loader"
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/http/v1/router"
	"github.com/NoobforAl/Enpass/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	validator := validation.New()
	store, err := database.New(env.GetDSN())
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.Default(r, store, validator)

	api := r.Group("/api")
	router.MainApi(api, store, validator)

	if err := r.Run("127.0.0.1:1111"); err != nil {
		panic(err)
	}
}
