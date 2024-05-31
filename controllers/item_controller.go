package controllers

import (
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
}

type IItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &IItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, items)
}
