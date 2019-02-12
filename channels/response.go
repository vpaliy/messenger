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
	Channel   uint      `json:"channel"`
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

type userSubscriptionsResponse struct {
	User          participant             `json:"user"`
	Subscriptions []*subscriptionResponse `json:"subscriptions"`
}

func newParticipant(u *model.User) participant {
	return participant{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
		Image:    u.Image,
	}
}

func newSubscriptionResponse(s *model.Subscription) *subscriptionResponse {
	return &subscriptionResponse{
		ID:        s.ID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		Private:   s.Private,
		Snippet:   s.Snippet,
		Channel:   s.ChannelID,
		Unread:    s.Unread,
	}
}

func newUserSubscriptionsResponse(u *model.User, subs *model.Subscription) *userSubscriptionsResponse {
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
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		Private:     c.Private,
		Description: c.Description,
		Tags:        c.Tags,
		Image:       c.Image,
		Archived:    c.Archived,
		Members:     member,
		Creator:     newParticipant(c.User),
	}
}