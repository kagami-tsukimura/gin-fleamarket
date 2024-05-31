package repositories

import "gin-fleamarket/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}
