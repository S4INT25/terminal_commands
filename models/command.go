package models

import "gorm.io/gorm"

type Platform int

const (
	Linux   Platform = 1
	Windows Platform = 2
	Mac     Platform = 3
)

type Command struct {
	gorm.Model
	Name        string
	Description string
	Platform    Platform
	User        User
	UserID      uint
}
