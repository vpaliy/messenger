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

// represents a user subscription to a channel
type subscriptionResponse struct {
	ID        uint            `json:"id"`
	CreatedAt utils.Timestamp `json:"created_at"`
	UpdatedAt utils.Timestamp `json:"updated_at"`
	Private   bool            `json:"private"`
	Snippet   string          `json:"snippet"`
	Channel   uint            `json:"channel"`
	Unread    int16           `json:"unread"`
}

// represents a channel info
type channelResponse struct {
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

type userSubscriptionsResponse struct {
	User          participant             `json:"user"`
	Subscriptions []*subscriptionResponse `json:"subscriptions"`
}

type userChannelsResponse struct {
	User     participant `json:"user"`
	Channels []*channel  `json:"channels"`
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

func newSubscriptionResponse(s *model.Subscription) *subscriptionResponse {
	return &subscriptionResponse{
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

func newUserChannelsResponse(u *model.User, cs []*model.Channel) *userChannelsResponse {
	channels := make([]*channel, len(cs))
	for i, channel := range cs {
		channels[i] = newChannel(channel)
	}
	return &userChannelsResponse{
		User:     newParticipant(u),
		Channels: channels,
	}
}

func newUserSubscriptionsResponse(u *model.User, subs []*model.Subscription) *userSubscriptionsResponse {
	subscriptions := make([]*subscriptionResponse, len(subs))
	for i, s := range subs {
		subscriptions[i] = newSubscriptionResponse(s)
	}
	return &userSubscriptionsResponse{
		User:          newParticipant(u),
		Subscriptions: subscriptions,
	}
}

func newChannelResponse(c *model.Channel) *channelResponse {
	users := c.GetUsers()
	members := make([]participant, len(users))
	for i, member := range users {
		members[i] = newParticipant(member)
	}
	return &channelResponse{
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
