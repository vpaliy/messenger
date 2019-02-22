package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) fetchChannel(ch string, c echo.Context) (*model.Channel, error) {
	channel, err := h.channelStore.Fetch(ch)
	if err != nil {
		return nil, c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if channel == nil {
		return nil, c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return channel, nil
}

func (h *Handler) GetMessages(c echo.Context) error {
	request := new(FetchMessagesRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// fetch the channel by name or ID
	channel, err := h.fetchChannel(request.Channel, c)
	// if there is an error, stop
	if channel == nil {
		return err
	}
	// check if the user is subscribed to the channel
	if !channel.HasUser(utils.GetUser(c).ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	// get all the messages
	messages, err := h.messageStore.GetForChannel(
		request.Channel,
		store.From(request.Oldest.Time()),
		store.To(request.Latest.Time()),
		store.Limit(request.Limit),
	)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewFetchMessagesResponse(channel, messages))
}

func (h *Handler) PostMessage(c echo.Context) error {
	request := new(CreateMessageRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.fetchChannel(request.Channel, c)
	// if there is an error
	if channel == nil {
		return err
	}
	// check if the user is subscribed to the chat
	currentUser := utils.GetUser(c)
	if !channel.HasUser(currentUser.ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	// submit a message
	message := request.ToMessage(
		currentUser.ID, channel.ID,
	)
	// create message and send an error message if fails
	if err := h.messageStore.Create(message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, NewCreateMessageResponse(message))
}

func (h *Handler) Search(c echo.Context) error {
	request := new(SearchMessagesRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.fetchChannel(request.Channel, c)
	// if there is an error or the channel doesn't exist, stop
	if channel == nil {
		return err
	}
	// search for messages
	messages, err := h.messageStore.Search(
		request.Channel,
		request.Query,
		store.From(request.Oldest.Time()),
		store.To(request.Latest.Time()),
		store.Limit(request.Limit),
	)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewFetchMessagesResponse(channel, messages))
}

func (h *Handler) DeleteMessage(c echo.Context) error {
	request := new(DeleteMessageRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	message, err := h.messageStore.Fetch(request.ID)
	// an error has occurred
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// message was not found
	if message == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	// only the author can delete the message
	// TODO: allow admins to delete the message too
	if message.UserID != utils.GetUser(c).ID {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	// delete it
	if err := h.messageStore.Delete(message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) EditMessage(c echo.Context) error {
	request := new(EditMessageRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	message, err := h.messageStore.Fetch(request.ID)
	// an error has occurred
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// message was not found
	if message == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	// only the author can edit the message
	// TODO: allow other people to edit the message too
	if message.UserID != utils.GetUser(c).ID {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	// update the message
	message.Text = request.Text
	if err := h.messageStore.Update(message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}
