package channels

import (
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
)

// represents a subscriber to a channel (or the creator of a channel)
type participant struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Fullname string  `json:"fullname"`
	Image    *string `json:"image"`
}

// represents a channel info
type channel struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	CreatedAt   utils.Timestamp `json:"created_at"`
	UpdatedAt   utils.Timestamp `json:"updated_at"`
	CreatorID   uint            `json:"creator_id"`
	Private     bool            `json:"private"`
	Description string          `json:"description"`
	Tags        []string        `json:"tags"`
	Image       *string         `json:"image"`
	Archived    bool            `json:"archived"`
	Members     []string        `json:"members"`
}

// represents a channel info
type ChannelResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	CreatedAt   utils.Timestamp `json:"created_at"`
	UpdatedAt   utils.Timestamp `json:"updated_at"`
	Private     bool            `json:"private"`
	Description string          `json:"description"`
	Tags        []string        `json:"tags"`
	Image       *string         `json:"image"`
	Archived    bool            `json:"archived"`
	Creator     participant     `json:"creator"`
	Members     []participant   `json:"members"`
}

// represents a user subscription to a channel
type SubscriptionResponse struct {
	ID        uint            `json:"id"`
	CreatedAt utils.Timestamp `json:"created_at"`
	UpdatedAt utils.Timestamp `json:"updated_at"`
	Private   bool            `json:"private"`
	Snippet   string          `json:"snippet"`
	Channel   uint            `json:"channel"`
	Unread    int16           `json:"unread"`
}

type UserSubscriptionsResponse struct {
	User          participant             `json:"user"`
	Subscriptions []*SubscriptionResponse `json:"subscriptions"`
}

type UserChannelsResponse struct {
	User     participant `json:"user"`
	Channels []*channel  `json:"channels"`
}

type ChannelsResponse struct {
	Channels []*channel `json:"channels"`
}

func newParticipant(u *model.User) participant {
	return participant{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.FullName,
		//Image:    u.Image,
	}
}

func NewSubscriptionResponse(s *model.Subscription) *SubscriptionResponse {
	return &SubscriptionResponse{
		ID:        s.ID,
		CreatedAt: utils.Timestamp(s.CreatedAt),
		UpdatedAt: utils.Timestamp(s.UpdatedAt),
		Private:   s.Private,
		Snippet:   s.Snippet,
		Channel:   s.ChannelID,
		Unread:    s.Unread,
	}
}

func newChannel(c *model.Channel) *channel {
	members := make([]string, len(c.Members))
	for i, member := range c.Members {
		members[i] = string(member.UserID)
	}
	return &channel{
		ID:          c.ID,
		Name:        c.Name,
		CreatorID:   c.CreatorID,
		CreatedAt:   utils.Timestamp(c.CreatedAt),
		UpdatedAt:   utils.Timestamp(c.UpdatedAt),
		Private:     c.Private,
		Description: c.Description,
		Tags:        c.GetTags(),
		//	Image:       c.Image,
		Archived: c.Archived,
		Members:  members,
	}
}

func NewChannelsResponse(cs []*model.Channel) *ChannelsResponse {
	channels := make([]*channel, len(cs))
	for i, channel := range cs {
		channels[i] = newChannel(channel)
	}
	return &ChannelsResponse{
		Channels: channels,
	}
}

func NewUserChannelsResponse(u *model.User, cs []*model.Channel) *UserChannelsResponse {
	channels := make([]*channel, len(cs))
	for i, channel := range cs {
		channels[i] = newChannel(channel)
	}
	return &UserChannelsResponse{
		User:     newParticipant(u),
		Channels: channels,
	}
}

func NewUserSubscriptionsResponse(u *model.User, subs []*model.Subscription) *UserSubscriptionsResponse {
	subscriptions := make([]*SubscriptionResponse, len(subs))
	for i, s := range subs {
		subscriptions[i] = NewSubscriptionResponse(s)
	}
	return &UserSubscriptionsResponse{
		User:          newParticipant(u),
		Subscriptions: subscriptions,
	}
}

func NewChannelResponse(c *model.Channel) *ChannelResponse {
	users := c.GetUsers()
	members := make([]participant, len(users))
	for i, member := range users {
		members[i] = newParticipant(member)
	}
	return &ChannelResponse{
		ID:          c.ID,
		Name:        c.Name,
		CreatedAt:   utils.Timestamp(c.CreatedAt),
		UpdatedAt:   utils.Timestamp(c.UpdatedAt),
		Private:     c.Private,
		Description: c.Description,
		Tags:        c.GetTags(),
		//	Image:       c.Image,
		Archived: c.Archived,
		Members:  members,
		Creator:  newParticipant(&c.Creator),
	}
}
