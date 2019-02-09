package model

const (
	Dialog  = "dialog"
	Group   = "group"
	General = "general"
)

type Channel struct {
	Title       string
	Tags        []string
	Creator     string
	Image       *string
	Description string
	Type        string
	Archived    bool
}

type Subcription struct {
	Alert     bool
	Unread    int16
	ChannelID string
	User      *User
	Type      string
}
