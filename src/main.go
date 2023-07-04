package main

import (
	env "github.com/NoobforAl/Enpass/config_loader"
	"github.com/NoobforAl/Enpass/controller"
	"github.com/NoobforAl/Enpass/database"
	"github.com/gin-gonic/gin"
)

func main() {
	stor, err := database.New(env.GetDSN())
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/genRandomPass", controller.GenRandomPass)
	r.POST("/login", controller.Login(stor))

	api := r.Group("/api")
	api.Use(controller.AuthMiddleware())
	{
		api.GET("/allPass", controller.AllPass(stor))
		api.GET("/allService", controller.AllService(stor))

		api.GET("/pass/:id", controller.FindPass(stor))
		api.GET("/service/:id", controller.FindService(stor))

		api.POST("/createPass", controller.NewPass(stor))
		api.POST("/createService", controller.NewService(stor))

		api.PUT("/updateUser", controller.UpdateUser(stor))
		api.PUT("/updatePass", controller.UpdatePass(stor))
		api.PUT("/updateService", controller.UpdateService(stor))

		api.DELETE("/deletePass/:id", controller.DeletePassWord(stor))
		api.DELETE("/deleteService/:id", controller.DeleteService(stor))
	}

	if err := r.Run("127.0.0.1:1111"); err != nil {
		panic(err)
	}
}
