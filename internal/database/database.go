package database

import "gorm.io/gorm"

type UserData interface {
	AddUser(userId int64)
}

type DataBase struct {
	UserData
}

func NewDataBase(db *gorm.DB) *DataBase {
	return &DataBase{
		UserData: nil,
	}
}
