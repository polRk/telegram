package telegram

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"`
	Selective       bool                `json:"selective,omitempty"`
}

// Validate returns an error if value is invalid.
func (k *ReplyKeyboardMarkup) Validate() error {
	for i := range k.Keyboard {
		for j := range k.Keyboard[i] {
			if err := k.Keyboard[i][j].Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
