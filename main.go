package main

import (
	"fmt"

	"github.com/Kunal-Patro/NoteTakingApp/controllers"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.MigrateDatabase()
}

func main() {
	fmt.Println("Getting Started...")

	router := gin.Default()

	router.POST("/signup", controllers.SignUp)

	router.POST("/login", controllers.Login)

	router.GET("/validate", middleware.ProcessAuth, controllers.Validate)

	router.Run()
}
