package api

import (
	"backend/internal/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, io websocket.IO) {
	api := r.Group("/api")
	api.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Hello, world!",
		})
	})

	api.GET("/broadcast", func(c *gin.Context) {
		io.Broadcast([]byte("Hello, world!"))
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Broadcasted!",
		})
	})
}
