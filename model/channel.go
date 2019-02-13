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
	CreatorID   uint
	Image       string
	Description string `gorm:"size:2048"`
	Type        string
	Archived    bool
	Members     []Subscription
	Private     bool
	Tags        []Tag `gorm:"many2many:channel_tags;association_autocreate:false"`
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

type Tag struct {
	gorm.Model
	Tag      string    `gorm:"unique_index"`
	Channels []Channel `gorm:"many2many:channel_tags;"`
}

func CreateTags(tags []string) []Tag {
	result := make([]Tag, len(tags))
	for i, t := range tags {
		result[i].Tag = t
	}
	return result
}

func (c *Channel) GetUsers() []*User {
	users := make([]*User, len(c.Members))
	for i, member := range c.Members {
		users[i] = &member.User
	}
	return users
}

func (c *Channel) GetTags() []string {
	tags := make([]string, len(c.Tags))
	for i, tag := range c.Tags {
		tags[i] = tag.Tag
	}
	return tags
}

func (c *Channel) IsCreator(id uint) bool {
	return c.CreatorID == id
}

func (c *Channel) HasUser(id uint) bool {
	for _, sub := range c.Members {
		if sub.UserID == id {
			return true
		}
	}
	return false
}

func (c *Channel) CreateSubscription(id uint) *Subscription {
	return &Subscription{
		ChannelID: c.ID,
		Private:   c.Private,
		UserID:    id,
	}
}
