package service

import "github.com/go-bot/internal/database"

type UserServiceImpl struct {
	userData *database.DataBase
}

func NewUserServiceImpl(database *database.DataBase) *UserServiceImpl {
	return &UserServiceImpl{userData: database}
}

func (u *UserServiceImpl) AddUser(userId int64) {
	u.userData.AddUser(userId)
}
