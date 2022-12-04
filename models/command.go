package models

import "gorm.io/gorm"

type Platform int

const (
	Linux   Platform = 0
	Windows Platform = 1
	Mac     Platform = 2
)

type Command struct {
	gorm.Model
	Name        string
	Description string
	Platform    Platform
	User        User
	UserID      uint
}

type CommandResponse struct {
	Id          uint
	Name        string
	Description string
	Platform    string
	User        UserResponse
}

func (p Platform) Name() string {
	return [...]string{"Linux", "Windows", "Mac"}[p]
}

func (c Command) ToResponse() CommandResponse {
	return CommandResponse{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Platform:    c.Platform.Name(),
		User: UserResponse{
			Id:    c.User.ID,
			Name:  c.User.Name,
			Email: c.User.Email,
		},
	}
}
