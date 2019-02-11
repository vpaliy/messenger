package messages

type attachment struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
	AudioURL string `json:"audio_url"`
	VideoURL string `json:"video_url"`
	ThumbURL string `json:"thumb_url"`
	Color    string `json:"color"`
}

type createMessageRequest struct {
	Channel     string       `json:"channel" validate:"required"`
	Text        string       `json:"text"`
	Attachments []attachment `json:"attachments"`
}
