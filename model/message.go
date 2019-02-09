package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Message struct {
		gorm.Model
		ChannelID   string
		User        *User
		Text        string
		Mentions    []*User
		Attachments []Attachment
		EditedAt    *time.Time
	}

	Attachment struct {
		Title    string
		Text     string
		ImageURL string
		AudioURL string
		VideoURL string
		ThumbURL string
		Color    string
	}
)
