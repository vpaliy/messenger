package rtm

type ResponseMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func NewResponse(event string, data interface{}) *ResponseMessage {
	return &ResponseMessage{event, data}
}
