package repositories

import "gorm.io/gorm"

type IAuthRepository interface {
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}
