package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/", Hello)
	r.GET("/health", Health)
	r.GET("/users", Users)
	r.GET("/products", Products)

	return r
}
