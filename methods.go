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

type SendMessagePayload struct {
	// ChatID is a unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID int64 `json:"chat_id"`

	// Text of the message to be sent, 1-4096 characters after entities parsing.
	Text string `json:"text"`

	// Mode for parsing entities in the message text.
	//
	// Optional.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Entities is a JSON-serialized list of special entities that appear in message text,
	// which can be specified instead of ParseMode.
	//
	// Optional.
	Entities []*MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview disables link previews for links in this message.
	//
	// Optional.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// DisableNotification sends the message silently.
	// Users will receive a notification with no sound.
	//
	// Optional.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ProtectContent protects the contents of the sent message from forwarding and saving.
	//
	// Optional.
	ProtectContent bool `json:"protect_content,omitempty"`

	// ReplyToMessageID is the ID of the original message.
	//
	// Optional.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// AllowSendingWithoutReply pass True, if the message should be sent even
	// if the specified replied-to message is not found.
	//
	// Optional.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// Additional interface options.
	// A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	// One of: InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply.
	//
	// Optional.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// SendMessage sending text messages.
// Returns sent Message on success.
func (c *Client) SendMessage(ctx context.Context, payload *SendMessagePayload) (*Message, error) {
	resp, err := c.MakeRequest(ctx, "sendMessage", payload)
	if err != nil {
		return nil, err
	}

	var message Message
	err = json.Unmarshal(resp.Result, &message)
	return &message, err
}
