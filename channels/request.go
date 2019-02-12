package channels

type createChannelRequest struct {
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Type        string   `json:"type" validate:"required"`
	Private     bool     `json:"private"`
}

type updateChannelRequest struct {
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Private     bool     `json:"private"`
}

func (r *updateChannelRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *createChannelRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (c *createChannelRequest) toChannel() *model.Channel {
	return &model.Channel{
		Name:        c.Channel,
		Tags:        c.Tags,
		Image:       c.Image,
		Description: c.Description,
		Type:        c.Type,
		Private:     c.Private,
	}
}

func (c *updateChannelRequest) toChannel() *model.Channel {
	return &model.Channel{
		Tags:        c.Tags,
		Image:       c.Image,
		Description: c.Description,
		Private:     c.Private,
	}
}
