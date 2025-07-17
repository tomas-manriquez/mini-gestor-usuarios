package routes

import (
	"go-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	//services.InitMongo()

	router.GET("/ping", controllers.PingController)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserById)
}
