package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})

	r.GET("/v2/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world v2",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
