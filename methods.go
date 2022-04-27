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

type AnswerWebAppQueryPayload struct {
	// WebAppQueryId is a unique identifier for the query to be answered.
	WebAppQueryId string `json:"web_app_query_id"`

	// Result is a JSON-serialized object describing the message to be sent.
	Result *InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery set the result of an interaction with a Web App and
// send a corresponding message on behalf of the user to the chat from which the query originated.
// On success, a SentWebAppMessage object is returned.
func (c *Client) AnswerWebAppQuery(ctx context.Context, payload *AnswerWebAppQueryPayload) (*SentWebAppMessage, error) {
	resp, err := c.MakeRequest(ctx, "answerWebAppQuery", payload)
	if err != nil {
		return nil, err
	}

	var webAppMessage SentWebAppMessage
	err = json.Unmarshal(resp.Result, &webAppMessage)
	return &webAppMessage, err
}

type SetChatMenuButtonPayload struct {
	// ChatID is a unique identifier for the target private chat.
	// If not specified, default bot's menu button will be changed.
	//
	// Optional.
	ChatID int64 `json:"chat_id,omitempty"`

	// MenuButton is a JSON-serialized object for the new bot's menu button.
	// Defaults to MenuButtonDefault.
	// Optional.
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButton change the bot's menu button in a ChatTypePrivate, or the default menu button.
// Returns True on success.
func (c *Client) SetChatMenuButton(ctx context.Context, payload *SetChatMenuButtonPayload) (bool, error) {
	resp, err := c.MakeRequest(ctx, "setChatMenuButton", payload)
	if err != nil {
		return false, err
	}

	var success bool
	err = json.Unmarshal(resp.Result, &success)
	return success, err
}

type GetChatMenuButtonPayload struct {
	// ChatID is a unique identifier for the target private chat.
	// If not specified, default bot's menu button will be returned.
	//
	// Optional.
	ChatID int64 `json:"chat_id,omitempty"`
}

// GetChatMenuButton get the current value of the bot's menu button in a private chat, or the default menu button.
// Returns MenuButton on success.
func (c *Client) GetChatMenuButton(ctx context.Context, payload *GetChatMenuButtonPayload) (*MenuButton, error) {
	resp, err := c.MakeRequest(ctx, "setChatMenuButton", payload)
	if err != nil {
		return nil, err
	}

	var menuButton MenuButton
	err = json.Unmarshal(resp.Result, &menuButton)
	return &menuButton, err
}

type SetMyDefaultAdministratorRightsPayload struct {
	// Rights is a JSON-serialized object describing new default administrator rights.
	// If not specified, the default administrator rights will be cleared.
	//
	// Optional.
	Rights *ChatAdministratorRights `json:"rights,omitempty"`

	// ForChannels is used to change the default administrator rights of the bot in channels.
	// Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	//
	// Optional.
	ForChannels bool `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights change the default administrator rights requested by the bot
// when it's added as an administrator to groups or channels.
// These rights will be suggested to users, but they are free to modify the list before adding the bot.
// Returns True on success.
func (c *Client) SetMyDefaultAdministratorRights(ctx context.Context, payload *SetMyDefaultAdministratorRightsPayload) (bool, error) {
	resp, err := c.MakeRequest(ctx, "setMyDefaultAdministratorRights", payload)
	if err != nil {
		return false, err
	}

	var success bool
	err = json.Unmarshal(resp.Result, &success)
	return success, err
}

type GetMyDefaultAdministratorRightsPayload struct {
	// ForChannels is used to change the default administrator rights of the bot in channels.
	// Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	//
	// Optional.
	ForChannels bool `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRights get the current default administrator rights of the bot.
// Returns ChatAdministratorRights on success.
func (c *Client) GetMyDefaultAdministratorRights(ctx context.Context, payload *GetMyDefaultAdministratorRightsPayload) (*ChatAdministratorRights, error) {
	resp, err := c.MakeRequest(ctx, "getMyDefaultAdministratorRights", payload)
	if err != nil {
		return nil, err
	}

	var chatAdministratorRights ChatAdministratorRights
	err = json.Unmarshal(resp.Result, &chatAdministratorRights)
	return &chatAdministratorRights, err
}
