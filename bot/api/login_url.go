// Bot API 5.7

package api

import (
	"errors"
)

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	// URL is an HTTP URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data,
	// the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	// NOTE: You must always check the hash of the received data
	// to verify the authentication and the integrity of the data as described in Checking authorization.
	URL string `json:"url"`

	// ForwardText is a new text of the button in forwarded messages.
	//
	// Optional.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername is a username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details.
	// If not specified, the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	//
	// Optional.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess pass True to request the permission for your bot to send messages to the user.
	//
	// Optional.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// Validate returns an error if value is invalid.
func (b *LoginURL) Validate() error {
	if len(b.URL) == 0 {
		return errors.New("missing required filed: URL")
	}

	return nil
}
