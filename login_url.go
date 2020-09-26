package telegram

import "fmt"

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// Validate returns an error if value is invalid.
func (b *LoginURL) Validate() error {
	if len(b.URL) == 0 {
		return fmt.Errorf("telegram[LoginURL]: URL is required")
	}

	return nil
}
