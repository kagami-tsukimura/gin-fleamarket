package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルーターを初期化
	r := gin.Default()
	// path, 無名関数
	r.GET("/sample", func(c *gin.Context) {
		// status_code, body
		// gin.H: map
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // default: 0.0.0.0:8080 でサーバーを立てます。
}
