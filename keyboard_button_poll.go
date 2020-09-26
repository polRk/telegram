package telegram

// KeyboardButtonPoll represents type of a poll, which is allowed
// to be created and sent when the corresponding button is pressed.
type KeyboardButtonPoll struct {
	Type *PollType `json:"type,omitempty"`
}

// Validate returns an error if value is invalid.
func (b *KeyboardButtonPoll) Validate() error {
	if b.Type != nil {
		if err := b.Type.Validate(); err != nil {
			return err
		}
	}

	return nil
}
