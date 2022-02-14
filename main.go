package main

import (
	"gin-mongo-api/configs"

	"github.com/gin-gonic/gin"

	"gin-mongo-api/routes"
)

func main() {
	router := gin.Default()

	routes.UserRoute(router)

	configs.COnnectToDB()

	router.Run("localhost:8080")
}
