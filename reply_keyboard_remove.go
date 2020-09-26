package telegram

// ReplyKeyboardRemove represents removing keyboard method.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// Validate returns an error if value is invalid.
func (k *ReplyKeyboardRemove) Validate() error {
	return nil
}
