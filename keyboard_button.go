package telegram

import "fmt"

// KeyboardButton represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object
// to specify text of the button.
type KeyboardButton struct {
	Text            string              `json:"text"`
	RequestContact  bool                `json:"request_contact,omitempty"`
	RequestLocation bool                `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPoll `json:"request_poll,omitempty"`
}

// Validate returns an error if value is invalid.
func (b *KeyboardButton) Validate() error {
	if len(b.Text) == 0 {
		return fmt.Errorf("telegram[KeyboardButton]: Text is required")
	}

	if b.RequestContact && b.RequestLocation || b.RequestLocation && b.RequestPoll != nil || b.RequestContact && b.RequestPoll != nil {
		return fmt.Errorf("telegram[KeyboardButton]: RequestContact, RequestLocation, and RequestPoll are mutually exclusive")
	}

	if b.RequestPoll != nil {
		if err := b.RequestPoll.Validate(); err != nil {
			return err
		}
	}

	return nil
}
