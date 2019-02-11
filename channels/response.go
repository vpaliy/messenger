package channels

import "time"

// represents a subscriber to a channel (or the creator of a channel)
type participant struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Fullname string  `json:"fullname"`
	Image    *string `json:"image"`
}

// represents a user subscription to a channel
type subscriptionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Private   bool      `json:"private"`
	Snippet   string    `json:"snippet"`
	ChannelID uint      `json:"channel_id"`
	Unread    int16     `json:"unread"`
}

// represents a channel info
type channelResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Private     bool          `json:"private"`
	Description string        `json:"description"`
	Tags        []string      `json:"tags"`
	Image       *string       `json:"image"`
	Archived    bool          `json:"archived"`
	Creator     participant   `json:"creator"`
	Members     []participant `json:"members"`
}
