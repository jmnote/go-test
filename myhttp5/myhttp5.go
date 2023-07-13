package myhttp5

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var limit = ratelimit.New(3)

func rateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit.Take()
	}
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.Use(rateLimit())
	api.GET("/ping", apiPing)
	return r
}

func apiPing(c *gin.Context) {
	c.String(200, "pong")
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}
