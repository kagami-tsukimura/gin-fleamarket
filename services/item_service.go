package services

import "gin-fleamarket/models"

type IItemService interface {
	FindAll() (*[]models.Item, error)
}
