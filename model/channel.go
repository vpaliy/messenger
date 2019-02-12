package model

import (
	"github.com/jinzhu/gorm"
)

const (
	Dialog  = "dialog"
	Group   = "group"
	General = "general"
)

type Channel struct {
	gorm.Model
	Name        string `gorm:"unique_index;not null"`
	Creator     User   `gorm:"foreignkey:CreatorID"`
	Tags        []string
	CreatorID   uint
	Image       string
	Description string `gorm:"size:2048"`
	Type        string
	Archived    bool
	Members     []Subscription
	Private     bool
}

type Subscription struct {
	gorm.Model
	Snippet   string
	Unread    int16
	Channel   Channel
	ChannelID uint
	User      User
	UserID    uint
	Type      string
	Private   bool
}

func (c *Channel) GetUsers() []*User {
	users := make([]*User, len(c.Members))
	for i, member := range c.Members {
		users[i] = &member.User
	}
	return users
}
