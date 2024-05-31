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
