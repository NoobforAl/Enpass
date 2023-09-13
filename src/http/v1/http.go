package http

import (
	env "github.com/NoobforAl/Enpass/config_loader"
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/http/v1/controller"
	"github.com/NoobforAl/Enpass/http/v1/router"
	"github.com/NoobforAl/Enpass/lib/caching"
	"github.com/NoobforAl/Enpass/loggers"
	"github.com/NoobforAl/Enpass/validation"
	"github.com/gin-gonic/gin"
)

func HttpApp() *gin.Engine {
	logger := loggers.New()
	env.EnvInit(logger)

	validator := validation.New(logger)
	cache := caching.New(env.GetLifeTime())
	store, err := database.New(env.GetDSN(), logger)
	if err != nil {
		logger.Panic(err)
	}

	gin.SetMode(env.GetGinMode())

	r := gin.Default()
	baseConfig := &controller.BaseConfig{
		GinApp:     r,
		Stor:       store,
		Logger:     logger,
		Cache:      cache,
		Validation: validator,
	}

	router.Default(baseConfig)
	router.MainApi(baseConfig)
	return r
}
