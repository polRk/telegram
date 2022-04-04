// Bot API 5.7

package api

import (
	"errors"
	"fmt"
)

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text is a label text on the button.
	Text string `json:"text"`
	// URL is HTTP or tg:// url to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username,
	// if this is allowed by their privacy settings.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// LoginURL is An HTTP URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	//
	// Optional.
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// CallbackData is data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	//
	// Optional.
	CallbackData string `json:"callback_data,omitempty"`

	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	//
	// Optional.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat if set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	//
	// Optional.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// CallbackGame is a description of the game that will be launched when the user presses the button.
	//
	// Optional.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay is a flag if True, send a Pay button.
	//
	// NOTE: This type of button must always be the first button in the first row
	// and can only be used in invoice messages.
	//
	// Optional.
	Pay bool `json:"pay,omitempty"`
}

// Validate returns an error if value is invalid.
func (b InlineKeyboardButton) Validate() error {
	optionalCount := 0

	if len(b.Text) == 0 {
		return errors.New("missing required field: Text")
	}

	if len(b.URL) != 0 {
		optionalCount += 1
	}

	if b.LoginURL != nil {
		optionalCount += 1

		if err := b.LoginURL.Validate(); err != nil {
			return err
		}
	}

	if len(b.CallbackData) != 0 {
		if len(b.CallbackData) > 64 {
			return errors.New("CallbackData must contains less then 64 bytes")
		}

		optionalCount += 1
	}

	if len(b.SwitchInlineQuery) != 0 {
		optionalCount += 1
	}

	if len(b.SwitchInlineQueryCurrentChat) != 0 {
		optionalCount += 1
	}

	if b.CallbackGame != nil {
		optionalCount += 1
	}

	if b.Pay != false {
		optionalCount += 1
	}

	if optionalCount > 1 {
		return fmt.Errorf("telegram[InlineKeyboardButton]: you must use exactly one of the optional fields")
	}

	return nil
}
