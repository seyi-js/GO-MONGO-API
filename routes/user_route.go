package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PATCH("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())

}
