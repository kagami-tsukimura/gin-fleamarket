package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemService struct {
	// Interface型で実装の切り替えを容易にする
	repository repositories.IItemRepository
}
