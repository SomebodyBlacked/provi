package routes

import "github.com/gin-gonic/gin"

func Products(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "products",
	})
}
