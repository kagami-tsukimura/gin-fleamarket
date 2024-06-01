package services

import "gin-fleamarket/repositories"

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(rpository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: rpository}
}
