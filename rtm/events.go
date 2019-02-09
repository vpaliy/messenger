package rtm

// this describes an interface between RTM API and the underlying database

type JoinParams struct {
	Channel string
}

type LeaveParams struct {
	Channel string
}

type SendParams struct {
	From    string
	Channel string
	Content interface{}
}

type LoadParams struct {
	From    string
	Channel string
	Limit   int
}

// an interface for interacting with the database
type EventHandler struct {
	// notify about joining
	OnJoin func(JoinParams) error

	// notify about unsubscribing
	OnLeave func(LeaveParams) error

	// notify about the recent message
	OnSend func(SendParams) error

	// request data from the database
	OnLoad func(LoadParams) ([]byte, error)
}
