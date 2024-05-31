package controllers

import (
	"gin-fleamarket/services"
	"net/http"

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
		// http.StatusInternalServerError: 500
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		// early return
		return
	}
	// http.StatusOK: 200
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
