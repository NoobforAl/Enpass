package main

import (
	env "github.com/NoobforAl/Enpass/config_loader"
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/http/v1/router"
	"github.com/NoobforAl/Enpass/loggers"
	"github.com/NoobforAl/Enpass/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := loggers.New()
	validator := validation.New(logger)
	store, err := database.New(env.GetDSN(), logger)
	if err != nil {
		logger.Panic(err)
	}

	r := gin.Default()
	router.Default(r, store, validator, logger)

	api := r.Group("/api")
	router.MainApi(api, store, validator, logger)

	if err := r.Run("0.0.0.0:1111"); err != nil {
		logger.Panic(err)
	}
}
