package controllers

import "github.com/gin-gonic/gin"

type IItemController interface {
	FindAll(ctx *gin.Context)
}
