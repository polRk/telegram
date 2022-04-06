package telegram

import (
	"context"
	"encoding/json"
)

// GetMe returns basic information about the bot in form of a User object.
func (c *Client) GetMe(ctx context.Context) (*User, error) {
	resp, err := c.MakeRequest(ctx, "getMe", nil)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(resp.Result, &user)
	return &user, err
}

// SendMessage sending text messages.
// Returns sent Message on success.
func (c *Client) SendMessage(ctx context.Context, msg *Message) (*Message, error) {
	resp, err := c.MakeRequest(ctx, "sendMessage", msg)
	if err != nil {
		return nil, err
	}

	var message Message
	err = json.Unmarshal(resp.Result, &message)
	return &message, err
}
