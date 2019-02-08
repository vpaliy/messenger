package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Message struct {
		gorm.Model
		RoomID      string
		User        *User
		Text        string
		Mentions    []string
		Attachments []Attachment
	}

	Attachment struct {
		Title    string
		ImageURL string
		AudioURL string
		VideoURL string
	}
)
