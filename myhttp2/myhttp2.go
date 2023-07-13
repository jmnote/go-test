package myhttp2

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}
	return router
}

func loginEndpoint(c *gin.Context) {
	c.String(200, "pong")
}

func submitEndpoint(c *gin.Context) {
	c.String(200, "pong")
}

func readEndpoint(c *gin.Context) {
	c.String(200, "pong")
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}
