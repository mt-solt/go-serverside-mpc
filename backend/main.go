package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// ルーター設定
	r := gin.Default()

	// CORS設定
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ルーティングの設定
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}
