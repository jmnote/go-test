package myhttp3b

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/v1/login", loginEndpoint)
	r.GET("/v1/submit", submitEndpoint)
	r.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.String(200, "pong")
}

func submitEndpoint(c *gin.Context) {
	c.String(200, "pong")
}
