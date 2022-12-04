package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	PassWordHash string
	Commands     []Command
}

type UserResponse struct {
	Id    uint
	Name  string
	Email string
}
