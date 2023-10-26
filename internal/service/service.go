package service

import (
	"github.com/go-bot/internal/database"
)

type UserService interface {
	AddUser(userId int64)
}

type Service struct {
	UserService
}

func NewService(database *database.DataBase) *Service {
	return &Service{
		UserService: nil,
	}
}
