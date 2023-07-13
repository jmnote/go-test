package myhttp6

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limit = rate.NewLimiter(rate.Every(time.Second), 3)

func rateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.String(429, "Too Many Requests")
			c.Abort()
		}
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
