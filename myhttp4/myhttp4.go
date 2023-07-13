package myhttp4

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/login", apiLogin)
	return r
}

func apiLogin(c *gin.Context) {
	// TODO: struct credential, apiError, func validate
	if c.PostForm("username") != "hello" || c.PostForm("password") != "world" {
		c.JSON(403, gin.H{"status": "error", "error": "username or password is incorrect"})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": gin.H{"token": "EXAMPLE", "message": "logged in successfully"}})
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}
