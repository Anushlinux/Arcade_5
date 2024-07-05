package main

import (
	"fmt"
	"gopro/controllers"
	"gopro/initializers"
	"gopro/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initializing...")
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDB()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}