package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	ItemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(ItemService)

	// Ginのルーターを初期化
	r := gin.Default()
	// path, 無名関数
	r.GET("/ping", func(c *gin.Context) {
		// status_code, body
		// gin.H: map
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.Run("localhost:8080") // default: 0.0.0.0:8080 でサーバーを立てます。
}
