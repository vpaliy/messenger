package rtm

type ResponseMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
