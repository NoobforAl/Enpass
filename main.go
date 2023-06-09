package main

import (
	"github.com/NoobforAl/Enpass/Db"
	"github.com/NoobforAl/Enpass/controller"
	env "github.com/NoobforAl/Enpass/loadEnv"
	"github.com/NoobforAl/Enpass/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	if _, err := Db.InitDB(env.GetDSN()); err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(logger.NewLogger())

	r.POST("/login", controller.Login)
	r.GET("/genRandomPass", controller.GenRandomPass)

	api := r.Group("/api")
	api.Use(controller.AuthMiddleware())
	{
		api.GET("/pass/:id", controller.FindPass)
		api.GET("/service/:id", controller.FindService)

		api.PUT("/updatePass", controller.UpdatePass)
		api.PUT("/updateService", controller.UpdateService)
		api.PUT("/updateUserPass", controller.UpdateUserPass)

		api.POST("/createPass", controller.NewPass)
		api.POST("/createService", controller.NewService)

		api.DELETE("/deletePass/:id", controller.DeletePassWord)
		api.DELETE("/deleteService/:id", controller.DeleteService)
	}

	if err := r.Run("127.0.0.1:1111"); err != nil {
		panic(err)
	}
}
