package main

import (
	env "github.com/NoobforAl/Enpass/config_loader"
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/router"
	"github.com/gin-gonic/gin"
)

func main() {
	stor, err := database.New(env.GetDSN())
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.Default(r, stor)

	api := r.Group("/api")
	router.MainApi(api, stor)

	if err := r.Run("127.0.0.1:1111"); err != nil {
		panic(err)
	}
}
