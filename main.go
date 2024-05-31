package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルーターを初期化
	r := gin.Default()
	// path, 無名関数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // default: 0.0.0.0:8080 でサーバーを立てます。
}
