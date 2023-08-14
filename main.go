package main

import (
	"fmt"

	"github.com/Kunal-Patro/NoteTakingApp/initializers"
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
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.Run()
}
