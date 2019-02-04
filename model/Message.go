package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Message struct {
		gorm.Model
		RoomID    uint
		User      *User
		Text      string
		ParseUrls bool
		Alias     string
		Mentions  []User
	}

	Attachment struct {
		Fields []AttachmentField
	}

	AttachmentField struct {
		Title string
		Value string
	}
)
