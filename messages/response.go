package messages

type message struct {
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	CreatedAt   time.Time    `json:"timestamp"`
	Attachments []attachment `json:"attachments"`
}

type createMessageResponse struct {
	Channel string  `json:"channel"`
	Message message `json:"message"`
}
