package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "users",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "products",
		})
	})
	return r
}
