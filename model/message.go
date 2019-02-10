package model

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Channel     Channel
	ChannelID   uint
	User        User
	UserID      uint
	Text        string
	Mentions    []User
	Attachments []Attachment
}

type Attachment struct {
	gorm.Model
	MessageID uint
	Title     string
	Text      string
	ImageURL  string
	AudioURL  string
	VideoURL  string
	ThumbURL  string
	Color     string
}
